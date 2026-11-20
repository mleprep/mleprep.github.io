"""Static HTML for every route of the site, plus sitemap.xml, robots.txt and 404.html (called by build.py).

Each page carries its own title, description, canonical URL, Open Graph tags and JSON-LD, and the same markup
assets/app.js renders, so crawlers that don't run JavaScript see the full content and browsers paint it before the
app takes over (the app re-renders in place). Markdown is rendered by md_html() below, a small renderer for the
subset content/ uses (paragraphs, lists, tables, blockquotes, fenced code, inline code/bold/italic/links); in the
browser the app renders the same text with marked.
"""
import hashlib
import html
import json
import math
import re
import shutil
from datetime import datetime, timezone
from pathlib import Path
from urllib.parse import urlencode

ORIGIN = "https://mleprep.github.io"
SITE = "SDE & MLE Prep"
OG_IMAGE = ORIGIN + "/assets/og.png"
DIFFS = ["Easy", "Medium", "Hard"]
BIG = ["Google", "Meta", "Amazon", "Microsoft", "Apple", "Bloomberg"]
# top-level folders that are not routes; a section id must never be one of these (its pages are rebuilt from scratch)
RESERVED = {"assets", "content", "data", "go", "java", "python", "tools", "browse", ".git", "__pycache__"}


def esc(s):
    return html.escape(str(s), quote=True)


# ---------------------------------------------------------------- markdown
FENCE_OPEN = re.compile(r"^(\s*)(`{3,})([\w-]*)\s*$")
LIST_ITEM = re.compile(r"^(\s*)([-*+]|\d+\.)\s+(.*)$")
TABLE_SEP = re.compile(r"^\s*\|?\s*:?-{2,}:?\s*(\|\s*:?-{2,}:?\s*)*\|?\s*$")


def inline(s):
    spans = []   # code spans become placeholders first, so nothing inside them is formatted

    def hold(m):
        spans.append(f"<code>{esc(m.group(2).strip() or m.group(2))}</code>")
        return f"\x00{len(spans) - 1}\x00"
    s = _fmt(re.sub(r"(`+)(.+?)\1", hold, s))
    return re.sub(r"\x00(\d+)\x00", lambda m: spans[int(m.group(1))], s)


def _fmt(s):
    s = esc(s)
    s = re.sub(r"\[([^\]]+)\]\(([^)\s]+)\)", lambda m: f'<a href="{m.group(2)}">{m.group(1)}</a>', s)
    s = re.sub(r"\*\*(?=\S)(.+?)(?<=\S)\*\*", r"<strong>\1</strong>", s)
    s = re.sub(r"(?<![\w*])\*(?=[^\s*])(.+?)(?<=[^\s*])\*(?![\w*])", r"<em>\1</em>", s)
    s = re.sub(r"(?<!\w)_(?=\S)(.+?)(?<=\S)_(?!\w)", r"<em>\1</em>", s)
    return s


def _cells(row):
    row = row.strip()
    if row.startswith("|"):
        row = row[1:]
    if row.endswith("|") and not row.endswith("\\|"):
        row = row[:-1]
    cells, cur = [], ""
    for ch in row:   # like GFM, a pipe splits cells even inside `code`; write \| to keep one
        if ch == "|" and not cur.endswith("\\"):
            cells.append(cur)
            cur = ""
        else:
            cur += ch
    return [c.strip().replace("\\|", "|") for c in cells + [cur]]


def _starts_block(lines, i):
    ln = lines[i]
    return (FENCE_OPEN.match(ln) or LIST_ITEM.match(ln) or ln.lstrip().startswith(">")
            or re.match(r"^#{1,6}\s", ln) or (ln.lstrip().startswith("|") and i + 1 < len(lines)
                                             and TABLE_SEP.match(lines[i + 1])))


def _indent(ln):
    return len(ln) - len(ln.lstrip(" "))


def blocks(lines):
    out, i, n = [], 0, len(lines)
    while i < n:
        ln = lines[i]
        if not ln.strip():
            i += 1
            continue
        m = FENCE_OPEN.match(ln)
        if m:
            pad, ticks, lang = m.groups()
            body, i = [], i + 1
            while i < n and not re.match(rf"^\s*{ticks}`*\s*$", lines[i]):
                body.append(lines[i][min(len(pad), _indent(lines[i])):])
                i += 1
            i += 1
            code = esc("\n".join(body) + "\n")
            cls = f' class="language-{lang}"' if lang else ""
            pre = '<pre class="viz">' if lang == "viz" else "<pre>"
            out.append(f"{pre}<code{cls}>{code}</code></pre>")
            continue
        h = re.match(r"^(#{1,6})\s+(.*?)\s*#*\s*$", ln)
        if h:
            out.append(f"<h{len(h.group(1))}>{inline(h.group(2))}</h{len(h.group(1))}>")
            i += 1
            continue
        if ln.lstrip().startswith("|") and i + 1 < n and TABLE_SEP.match(lines[i + 1]):
            head, aligns = _cells(ln), []
            for c in _cells(lines[i + 1]):
                aligns.append("center" if c.startswith(":") and c.endswith(":") else "right" if c.endswith(":")
                              else "left" if c.startswith(":") else "")
            al = lambda k: f' align="{aligns[k]}"' if k < len(aligns) and aligns[k] else ""
            rows, i = [], i + 2
            while i < n and lines[i].strip().startswith("|"):
                rows.append(_cells(lines[i]))
                i += 1
            t = "<table><thead><tr>" + "".join(f"<th{al(k)}>{inline(c)}</th>" for k, c in enumerate(head)) + "</tr></thead>"
            if rows:
                t += "<tbody>" + "".join("<tr>" + "".join(f"<td{al(k)}>{inline(c)}</td>" for k, c in enumerate(r[:len(head)]))
                                         + "</tr>" for r in rows) + "</tbody>"
            out.append(t + "</table>")
            continue
        if ln.lstrip().startswith(">"):
            q = []
            while i < n and lines[i].strip() and (lines[i].lstrip().startswith(">") or q):
                q.append(re.sub(r"^\s*>\s?", "", lines[i]))
                i += 1
            out.append(f"<blockquote>{blocks(q)}</blockquote>")
            continue
        m = LIST_ITEM.match(ln)
        if m:
            i = _list(lines, i, out)
            continue
        para = [ln.strip()]
        i += 1
        while i < n and lines[i].strip() and not _starts_block(lines, i):
            para.append(lines[i].strip())
            i += 1
        out.append(f"<p>{inline(chr(10).join(para))}</p>")
    return "".join(out)


def _list(lines, i, out):
    m = LIST_ITEM.match(lines[i])
    base, ordered = _indent(lines[i]), m.group(2)[0].isdigit()
    start = int(m.group(2)[:-1]) if ordered else 1
    items, loose, n = [], False, len(lines)
    while i < n:
        m = LIST_ITEM.match(lines[i])
        if not m or _indent(lines[i]) > base + 1 or m.group(2)[0].isdigit() != ordered:
            break
        content = _indent(lines[i]) + len(m.group(2)) + 1
        body, i = [m.group(3)], i + 1
        while i < n:
            ln = lines[i]
            if not ln.strip():   # blank: the item goes on only if indented content follows
                j = i
                while j < n and not lines[j].strip():
                    j += 1
                if j < n and _indent(lines[j]) >= content:
                    body += [""] * (j - i)
                    loose = loose or not LIST_ITEM.match(lines[j])
                    i = j
                    continue
                if j < n and LIST_ITEM.match(lines[j]) and _indent(lines[j]) <= base + 1:
                    loose, i = True, j
                break
            if _indent(ln) >= content or (_indent(ln) > base and LIST_ITEM.match(ln)):
                body.append(ln[min(content, _indent(ln)):])
            elif not _starts_block(lines, i) and body[-1].strip():
                body.append(ln.strip())   # lazy continuation of the item's paragraph
            else:
                break
            i += 1
        items.append(body)
    tag = "ol" if ordered else "ul"
    attr = f' start="{start}"' if ordered and start != 1 else ""
    html_items = []
    for body in items:
        b = blocks(body)
        if not loose and b.startswith("<p>"):
            k = b.index("</p>")
            b = b[3:k] + b[k + 4:]
        html_items.append(f"<li>{b}</li>")
    out.append(f"<{tag}{attr}>{''.join(html_items)}</{tag}>")
    return i


def md_html(text):
    return f'<div class="md">{blocks((text or "").split(chr(10)))}</div>'


def plain(md, want=150):
    """Leading prose of a markdown text as plain text (for meta descriptions)."""
    text = re.sub(r"^\s*`{3,}.*?^\s*`{3,}\s*$", "", md or "", flags=re.M | re.S)
    out = []
    for block in re.split(r"\n\s*\n", text):
        rows = [_cells(ln) for ln in block.split("\n") if ln.strip().startswith("|")]
        if len(rows) > 2:   # a table: "first cell: second cell" per body row
            s = "; ".join(": ".join(c for c in r[:2] if c) for r in rows[2:])
        else:
            lines = [ln for ln in block.split("\n") if ln.strip() and not ln.strip().startswith("|")]
            if not lines:
                continue
            parts = [re.sub(r"^\s*(?:[-*+>]|\d+\.)\s+", "", ln).strip() for ln in lines]
            for k in range(len(parts) - 1):   # "**Two traversal families**" right above a list
                if LIST_ITEM.match(lines[k + 1]) and not LIST_ITEM.match(lines[k]) and parts[k].rstrip("*_")[-1:].isalnum():
                    parts[k] += ":"
            s = " ".join(parts)
        s = re.sub(r"`+([^`]*)`+", r"\1", s)
        s = re.sub(r"\[([^\]]+)\]\([^)]*\)", r"\1", s)
        s = re.sub(r"(\*\*|__|\*|(?<!\w)_|_(?!\w))", "", s)
        s = re.sub(r"\s+", " ", s).strip()
        if s and s[-1].isalnum():   # a bold lead-in line such as "**Two traversal families**"
            s += ":"
        out.append(s)
        if len(" ".join(out)) >= want:
            break
    return " ".join(out)


def clip(s, n=160):
    s = re.sub(r"\s+", " ", s).strip()
    if len(s) <= n:
        return s
    cut = s[:n - 1]
    cut = cut[:cut.rfind(" ")] if " " in cut[n // 2:] else cut
    return cut.rstrip(" ,;:.—-(") + "…"


# ---------------------------------------------------------------- urls + small helpers (mirror assets/app.js)
def href(*p):
    return "/" + "".join(f"{x}/" for x in p)


def browse(params=None):
    return "/browse/" + ("?" + urlencode(params) if params else "")


def tag_label(t):
    return "★ Blind 75" if t == "blind75" else "🔒 Premium" if t == "premium" else t


def tag_class(t):
    return " star" if t == "blind75" else " prem" if t == "premium" else ""


def fmt_f(v):
    return str(int(math.floor(v + 0.5))) if v >= 10 else f"{v:.1f}"


def fit(*options):
    """The first (most descriptive) title that fits in 70 characters, else the last (shortest)."""
    return next((x for x in options if len(x) <= 70), options[-1])


def page_title(kind, s=None, t=None, it=None, n=None):
    """<title> per page: keep in sync with pageTitle() in assets/app.js."""
    if kind == "home":
        return f"{SITE}: LeetCode, System Design & ML Interview Prep"
    if kind == "browse":
        return f"All {n['dsa']} Problems: Filter by Company, List & Pattern | {SITE}"
    if kind == "section":
        return {"dsa": f"Algorithms: {n['dsa']} LeetCode Problems by Pattern",
                "system-design": f"System Design Interview: {n.get('system-design')} Articles",
                "ml": f"ML & Deep Learning Interview: {n.get('ml')} Articles"}.get(s["id"], s["title"]) + f" | {SITE}"
    if kind == "topic":
        if s["id"] == "dsa":
            return fit(f"{t['title']}: LeetCode Problems & Solutions | {SITE}", f"{t['title']}: LeetCode Problems & Solutions", t["title"])
        return fit(f"{t['title']} – {s['title']} | {SITE}", f"{t['title']} – {s['title']}", t["title"])
    if s["id"] == "dsa":
        b = f"{it['id']}. {it['title']}"
        return fit(f"{b} – LeetCode Solution in Python, Java & Go", f"{b} – LeetCode Solution", b)
    kind = "ML Interview" if s["id"] == "ml" else "System Design"
    return fit(f"{it['title']} – {kind} | {SITE}", f"{it['title']} – {kind}", it["title"])


class Site:
    def __init__(self, sections, companies, lists):
        self.sections = sections   # [{"id", "title", "topics": [{"id", "title", "intro", "items": [...]}]}]
        self.co = companies
        self.lists = (lists or {}).get("lists", [])
        self.flat = [{"s": s["id"], "t": t["id"], "tTitle": t["title"], "it": it}
                     for s in sections for t in s["topics"] for it in t["items"]]
        self.by_id = {f["it"]["id"]: f for f in self.flat if f["s"] == "dsa"}
        self.n = {s["id"]: sum(len(t["items"]) for t in s["topics"]) for s in sections}
        self.co_by = {}
        if companies:
            names = companies["companies"]
            self.co_by = {slug: [{"c": names[i], "a": a, "r": r} for i, a, r in rows]
                          for slug, rows in companies["problems"].items()}

    def sec(self, sid):
        return next(s for s in self.sections if s["id"] == sid)

    # ---- chrome
    def nav(self, cur):
        return "".join(f'<a href="{href(s["id"])}" class="{"on" if s["id"] == cur else ""}">{esc(s["title"])}</a>'
                       for s in self.sections)

    def side(self, sid, tid=None, on_browse=False):
        s = self.sec(sid)
        out = f"<h4>{esc(s['title'])}</h4>"
        if s["id"] == "dsa":
            out += f'<a href="/browse/" class="all{" on" if on_browse else ""}"><span>All problems</span><small>{self.n["dsa"]}</small></a>'
        return out + "".join(f'<a href="{href(s["id"], t["id"])}" title="{esc(t["title"])}" class="{"on" if t["id"] == tid else ""}">'
                             f'<span>{esc(t["title"])}</span><small>{len(t["items"])}</small></a>' for t in s["topics"])

    # ---- lists
    def rows(self, lst, show_topic=False):
        dsa = any(f["s"] == "dsa" for f in lst)
        out = []
        for f in lst:
            it = f["it"]
            tags = "".join(f'<a href="{browse({"list": "blind75"} if t == "blind75" else {"tag": t})}" class="tag{tag_class(t)}" '
                           f'data-tag="{esc(t)}">{esc(tag_label(t))}</a>' for t in it.get("tags", []))
            cells = f'<td class="c"><input type="checkbox" data-k="{f["s"]}/{it["slug"]}" aria-label="done"></td>'
            if dsa:
                cells += f'<td class="n">{it.get("id", "")}</td>'
            cells += f'<td class="t"><a href="{href(f["s"], f["t"], it["slug"])}">{esc(it["title"])}</a><span class="tags">{tags}</span></td>'
            if show_topic:
                dt = f' data-topic="{f["t"]}"' if f["s"] == "dsa" else ""
                cells += f'<td class="tp"><a href="{href(f["s"], f["t"])}"{dt}>{esc(f["tTitle"])}</a></td>'
            if it.get("difficulty"):
                d = it["difficulty"]
                cells += f'<td class="d"><a href="{browse({"d": d})}" class="{d}" data-d="{d}">{d}</a></td>'
            elif dsa:
                cells += "<td></td>"
            out.append(f"<tr>{cells}</tr>")
        return f'<div class="list-wrap"><table class="list"><tbody>{"".join(out)}</tbody></table></div>'

    def topic_cards(self, s):
        out = []
        for i, t in enumerate(s["topics"]):
            by = [(d, sum(it.get("difficulty") == d for it in t["items"])) for d in DIFFS]
            by = [(d, k) for d, k in by if k]
            out.append(f'<a class="card" href="{href(s["id"], t["id"])}"><b>{f"{i + 1}. " if s["id"] == "dsa" else ""}{esc(t["title"])}</b>'
                       f'<span>0/{len(t["items"])} done{" · " + " ".join(f'<em class="{d}">{k}</em>' for d, k in by) if by else ""}</span>'
                       f'<div class="bar-prog"><i style="width:0%"></i></div></a>')
        return f'<div class="cards">{"".join(out)}</div>'

    def co_rows(self, it):
        return self.co_by.get(it["slug"], [])

    def co_val(self, it, c, win):
        r = next((x for x in self.co_rows(it) if x["c"] == c), None)
        return 0 if not r else r["r"] if win == "6m" else (r["a"] or r["r"])

    def asked_at(self, it):
        rows = sorted(self.co_rows(it), key=lambda x: (0 if x["c"] in BIG else 1 if x["r"] else 2, -(x["a"] or x["r"])))
        if not rows:
            return ""
        top, src = 12, self.co["source"]
        title = (f"Company data: {src['repo'].replace('https://github.com/', '')}{' (' + src['date'] + ')' if src.get('date') else ''}."
                 " Frequency is relative within each company (100 = that company's most-asked problem).")
        chips = "".join(
            f'<a class="co{" extra" if k >= top else ""}{" hot" if x["r"] else ""}" href="{browse({"co": x["c"], "sort": "freq"})}" '
            f'title="{esc(x["c"])}: frequency {fmt_f(x["a"])} all time{", " + fmt_f(x["r"]) + " in the last 6 months" if x["r"] else ""}">'
            f'{esc(x["c"])}<span>{fmt_f(x["a"] or x["r"])}</span></a>' for k, x in enumerate(rows))
        more = f'<button id="comore" class="co more">+{len(rows) - top} more</button>' if len(rows) > top else ""
        return f'<div class="asked"><span class="flabel" title="{esc(title)}">Asked at {len(rows)}</span>{chips}{more}</div>'

    # ---- views
    def view_home(self):
        dsa, total = self.sec("dsa"), self.n["dsa"]
        counts = {d: sum(f["it"].get("difficulty") == d for f in self.flat if f["s"] == "dsa") for d in DIFFS}
        has_go = sum("go" in f["it"] for f in self.flat if f["s"] == "dsa")
        b75 = sum("blind75" in f["it"].get("tags", []) for f in self.flat)
        langs = "Python, Java &amp; Go" if has_go >= total else f"Python &amp; Java{f' (and Go for {has_go} so far)' if has_go else ''}"
        h = f"""<h1>Interview prep, condensed.</h1>
    <p class="lead">{total} LeetCode problems — every one from Blind 75, NeetCode 150 &amp; 250, Grind 169, LeetCode 75 and Top Interview 150, plus
    frequently asked company problems — each with the key idea, a sketch, complexity and tested {langs}. Open any problem and
    pick <b>Solve in browser</b> to write Python{' or Go' if has_go else ''} and run it against the tests right on this page (nothing runs on a server).
    Plus system design and ML / deep learning, so one site covers both SDE and MLE interview loops. Progress is saved in this browser.</p>
    <div class="meta big">{"".join(f'<a class="badge {d}" href="{browse({"d": d})}">{counts[d]} {d}</a>' for d in DIFFS)}
      <a class="badge star" href="{browse({"list": "blind75"})}">★ {b75} Blind 75</a>
      <a href="{browse({"st": "done"})}">0/{total} done</a>
      <a href="{browse({"st": "todo"})}">Todo list →</a>
      <a href="#" id="rand">Random unsolved →</a></div>
    <div class="cards"><a class="card" href="/browse/"><b>Browse all problems</b><span>filter by difficulty, list, pattern, company, status</span></a>"""
        for s in self.sections:
            h += (f'<a class="card" href="{href(s["id"])}"><b>{esc(s["title"])}</b><span>{len(s["topics"])} topics · '
                  f'{self.n[s["id"]]} {"problems" if s["id"] == "dsa" else "articles"}</span></a>')
        h += "</div>"
        slugs = {f["it"]["slug"] for f in self.flat if f["s"] == "dsa"}
        if self.lists:
            h += "<h2>Study lists</h2><div class=\"cards\">" + "".join(
                f'<a class="card" href="{browse({"list": l["id"], "sort": "list"})}"><b>{"★ " if l["id"] == "blind75" else ""}{esc(l["name"])}</b>'
                f'<span>0/{len(slugs & set(l["slugs"]))} done · in the list\'s own order</span><div class="bar-prog"><i style="width:0%"></i></div></a>'
                for l in self.lists) + "</div>"
        if self.co_by:
            cards = []
            for c in BIG:
                every = [f for f in self.flat if f["s"] == "dsa" and self.co_val(f["it"], c, "")]
                recent = [f for f in every if self.co_val(f["it"], c, "6m")]
                cards.append(f'<a class="card" href="{browse({"co": c, "sort": "freq"})}"><b>{esc(c)}</b><span>{len(every)} problems · '
                             f'{len(recent)} in the last 6 months</span></a>')
            cards.append(f'<a class="card" href="{browse({"sort": "freq"})}"><b>Most asked overall</b><span>ranked by how many companies ask each problem</span></a>')
            h += f'<h2>By company</h2><div class="cards">{"".join(cards)}</div>'
        return h + f"<h2>Algorithm roadmap</h2>{self.topic_cards(dsa)}"

    def view_section(self, s):
        lead = '<p class="lead">Work through the topics in order, or <a href="/browse/">browse and filter all problems</a>.</p>' if s["id"] == "dsa" else ""
        return f"<h1>{esc(s['title'])}</h1>{lead}{self.topic_cards(s)}"

    def view_browse(self):
        lst = [f for f in self.flat if f["s"] == "dsa"]
        return (f'<div class="crumbs"><a href="{href("dsa")}">Algorithms</a></div><h1>All problems</h1><div id="lv">'
                f'<div class="fsum">Showing <b>{len(lst)}</b> of {len(lst)}</div>{self.rows(lst, show_topic=True)}</div><div id="arts"></div>')

    def view_topic(self, s, t):
        lst = [f for f in self.flat if f["s"] == s["id"] and f["t"] == t["id"]]
        intro = md_html(t["intro"]) if t.get("intro") else ""
        return (f'<div class="crumbs"><a href="{href(s["id"])}">{esc(s["title"])}</a></div><h1>{esc(t["title"])}</h1>'
                f'<div id="intro" class="article">{intro}</div><h2>{"Problems" if s["id"] == "dsa" else "Articles"}</h2>'
                f'<div id="lv"><div class="fsum">Showing <b>{len(lst)}</b> of {len(lst)}</div>{self.rows(lst)}</div>')

    def view_item(self, s, t, it):
        dsa = s["id"] == "dsa"
        pos = next(k for k, f in enumerate(self.flat) if f["s"] == s["id"] and f["t"] == t["id"] and f["it"]["slug"] == it["slug"])
        prev = self.flat[pos - 1] if pos > 0 and self.flat[pos - 1]["s"] == s["id"] else None
        nxt = self.flat[pos + 1] if pos + 1 < len(self.flat) and self.flat[pos + 1]["s"] == s["id"] else None
        tags = "".join(f'<a class="tag{tag_class(x)}" href="{browse({"list": "blind75"} if x == "blind75" else {"tag": x})}" '
                       f'title="Browse problems tagged {esc(x)}">{esc(tag_label(x))}</a>' for x in it.get("tags", []))
        lc = f"https://leetcode.com/problems/{it['slug']}/"
        meta = (f'<a class="badge {it["difficulty"]}" href="{browse({"d": it["difficulty"]})}" title="All {it["difficulty"]} problems">{it["difficulty"]}</a>'
                f'<a href="{lc}" target="_blank" rel="noopener">LeetCode ↗</a>' if dsa else "")
        h = (f'<div class="crumbs"><a href="{href(s["id"])}">{esc(s["title"])}</a> › <a href="{href(s["id"], t["id"])}">{esc(t["title"])}</a></div>'
             f'<h1>{str(it["id"]) + ". " if dsa else ""}{esc(it["title"])}</h1>'
             f'<div class="meta">{meta}{tags}<label class="done"><input type="checkbox" data-k="{s["id"]}/{it["slug"]}"> Done</label></div>')
        if dsa:
            related = [self.by_id[r] for r in it.get("related", []) if r in self.by_id]
            rel = ('<div class="related"><h3>Related problems</h3><div class="chips">' + "".join(
                f'<a href="{href(f["s"], f["t"], f["it"]["slug"])}"><i class="dot {f["it"]["difficulty"]}"></i>{f["it"]["id"]}. {esc(f["it"]["title"])}</a>'
                for f in related) + "</div></div>") if related else ""
            other = "".join(f'<pre hidden><code class="language-{lang}">{esc(it[lang])}</code></pre>' for lang in ("java", "go") if it.get(lang))
            h += (self.asked_at(it) +
                  f'<div class="split"><div class="desc"><div id="body">{md_html(it["md"])}</div>{rel}</div>'
                  '<div class="codecol"><div class="modes" role="tablist" aria-label="Mode"><button role="tab" data-mode="read" class="on" aria-selected="true">Solution</button>'
                  '<button role="tab" data-mode="solve" aria-selected="false">Solve in browser ▶</button></div>'
                  '<div class="code" id="readbox"><div class="bar"><button data-l="python" class="on">Python</button><button data-l="java">Java</button>'
                  + ('<button data-l="go">Go</button>' if it.get("go") else "") + '<span class="sp"></span>'
                  '<button id="tests">Tests</button><button id="copy">Copy</button><button id="file" title="Open the runnable file on GitHub">Run locally ↗</button></div>'
                  f'<div class="scroll"><pre><code id="src" class="language-python">{esc(it["py"])}</code></pre>'
                  f'<pre class="test" hidden><code id="tsrc">{esc(it["pyTest"])}</code></pre>{other}</div></div></div></div>')
        else:
            rail = "".join(f'<a href="{href(s["id"], t["id"], x["slug"])}" class="{"on" if x["slug"] == it["slug"] else ""}">{esc(x["title"])}</a>'
                           for x in t["items"])
            h += (f'<div class="artgrid"><div id="body" class="article">{md_html(it["md"])}</div>'
                  f'<nav class="rail"><h4>In {esc(t["title"])}</h4>{rail}</nav></div>')
        link = lambda f, txt: f'<a href="{href(f["s"], f["t"], f["it"]["slug"])}">{txt}</a>'
        return h + (f'<div class="pager"><span>{link(prev, "← " + esc(prev["it"]["title"])) if prev else ""}</span>'
                    f'<span>{link(nxt, esc(nxt["it"]["title"]) + " →") if nxt else ""}</span></div>')

    # ---- descriptions + structured data
    def desc_item(self, s, it):
        body = plain(it["md"])
        if s["id"] != "dsa":
            return clip(body if len(body) >= 100 else f"{it['title']}: {body}" if body else f"{it['title']}: {s['title']} interview notes.")
        d = clip(f"LeetCode {it['id']}. {it['title']} ({it['difficulty']}): {body}")
        return d if len(d) > 110 else clip(d + " Python, Java & Go solutions with tests you can run in the browser.")

    def desc_topic(self, s, t):
        intro = plain(t.get("intro", ""))
        if s["id"] == "dsa":
            return clip(f"{len(t['items'])} {t['title']} LeetCode problems with Python, Java & Go solutions. {intro}")
        covers = ", ".join(x["title"] for x in t["items"])
        return clip(f"{t['title']}: {intro}" + (f" Covers {covers}." if len(intro) < 100 else "") if intro
                    else f"{t['title']} ({s['title']}): {covers}")

    def desc_section(self, s):
        if s["id"] == "dsa":
            return clip(f"{self.n['dsa']} LeetCode problems grouped into {len(s['topics'])} patterns, from {s['topics'][0]['title']} to "
                        f"{s['topics'][-1]['title']}. Each has the key idea, complexity and tested Python, Java & Go solutions.")
        kind = "machine learning and deep learning" if s["id"] == "ml" else "system design"
        return clip(f"{self.n[s['id']]} {kind} interview articles: " + ", ".join(t["title"] for t in s["topics"]) + ".")


def crumbs_ld(trail):
    return {"@type": "BreadcrumbList", "itemListElement": [
        {"@type": "ListItem", "position": k + 1, "name": name, "item": ORIGIN + path} for k, (name, path) in enumerate(trail)]}


ORG = {"@type": "Organization", "name": SITE, "url": ORIGIN + "/"}


def page(shell, *, path, title, desc, main, nav, side, ld=None, og_type="website", noindex=False):
    url = ORIGIN + path
    head = [f'<link rel="canonical" href="{url}">'] if not noindex else ['<meta name="robots" content="noindex">']
    head += [f'<meta property="og:type" content="{og_type}">', f'<meta property="og:site_name" content="{esc(SITE)}">',
             f'<meta property="og:title" content="{esc(title)}">', f'<meta property="og:description" content="{esc(desc)}">',
             f'<meta property="og:url" content="{url}">', f'<meta property="og:image" content="{OG_IMAGE}">',
             '<meta property="og:image:width" content="1200">', '<meta property="og:image:height" content="630">',
             f'<meta property="og:image:alt" content="{esc(SITE)}: LeetCode, system design and ML interview prep">',
             '<meta name="twitter:card" content="summary_large_image">']
    if ld:
        data = json.dumps({"@context": "https://schema.org", "@graph": ld}, ensure_ascii=False, separators=(",", ":"))
        head.append(f'<script type="application/ld+json">{data.replace("</", "<" + chr(92) + "/")}</script>')
    return (shell.replace("{{title}}", esc(title)).replace("{{description}}", esc(desc)).replace("{{head}}", "\n".join(head))
            .replace("{{nav}}", nav).replace("{{side}}", side).replace("{{main}}", main))


def write_site(root, sections, companies, lists, state_file):
    root = Path(root)
    shell = (root / "tools" / "shell.html").read_text(encoding="utf-8")
    site = Site(sections, companies, lists)
    pages = []   # (path, html, hash-of-content)

    def add(path, **kw):
        h = page(shell, path=path, **kw)
        digest = hashlib.sha256((kw["title"] + kw["desc"] + kw["main"]).encode()).hexdigest()[:16]
        pages.append((path, h, digest, kw.get("noindex", False)))

    n = site.n
    add("/", title=page_title("home"), main=site.view_home(), nav=site.nav(""), side=site.side("dsa"),
        desc=clip(f"{n['dsa']} LeetCode problems from Blind 75, NeetCode 150/250, Grind 169 and more, with Python, Java & Go "
                  "solutions you can run in the browser, plus system design and ML / deep learning interview prep."),
        ld=[{"@type": "WebSite", "name": SITE, "alternateName": "mleprep", "url": ORIGIN + "/", "inLanguage": "en",
             "description": "SDE and MLE interview prep: algorithms with an in-browser judge, system design, ML and deep learning."}, ORG])
    add("/browse/", title=page_title("browse", n=n), main=site.view_browse(), nav=site.nav("dsa"), side=site.side("dsa", on_browse=True),
        desc=clip(f"Filter all {n['dsa']} LeetCode problems by difficulty, study list (Blind 75, NeetCode 150, Grind 169…), pattern and "
                  "company, with Python, Java & Go solutions."),
        ld=[crumbs_ld([("Algorithms", href("dsa")), ("All problems", "/browse/")])])
    for s in sections:
        add(href(s["id"]), title=page_title("section", s=s, n=n), desc=site.desc_section(s), main=site.view_section(s),
            nav=site.nav(s["id"]), side=site.side(s["id"]), ld=[crumbs_ld([(s["title"], href(s["id"]))])])
        for t in s["topics"]:
            trail = [(s["title"], href(s["id"])), (t["title"], href(s["id"], t["id"]))]
            add(href(s["id"], t["id"]), title=page_title("topic", s=s, t=t), desc=site.desc_topic(s, t), main=site.view_topic(s, t),
                nav=site.nav(s["id"]), side=site.side(s["id"], t["id"]), ld=[crumbs_ld(trail)])
            for it in t["items"]:
                path, desc = href(s["id"], t["id"], it["slug"]), site.desc_item(s, it)
                name = f"{it['id']}. {it['title']}" if s["id"] == "dsa" else it["title"]
                art = {"@type": "TechArticle", "headline": clip(page_title("item", s=s, it=it).replace(f" | {SITE}", ""), 110),
                       "name": name, "description": desc, "url": ORIGIN + path, "mainEntityOfPage": ORIGIN + path,
                       "inLanguage": "en", "isAccessibleForFree": True, "image": OG_IMAGE,
                       "license": "https://creativecommons.org/licenses/by/4.0/", "author": ORG, "publisher": ORG,
                       "isPartOf": {"@type": "WebSite", "name": SITE, "url": ORIGIN + "/"}}
                if it.get("tags"):
                    art["keywords"] = ", ".join(x for x in it["tags"] if x not in ("premium",))
                if s["id"] == "dsa":
                    art["about"] = f"LeetCode problem {it['id']}: {it['title']}"
                    art["educationalLevel"] = it["difficulty"]
                add(path, title=page_title("item", s=s, t=t, it=it), desc=desc, main=site.view_item(s, t, it),
                    nav=site.nav(s["id"]), side=site.side(s["id"], t["id"]), og_type="article",
                    ld=[crumbs_ld(trail + [(name, path)]), art])
    add("/404.html", title=f"Not found | {SITE}", desc="This page does not exist.", noindex=True,
        main='<h1>Not found</h1><p><a href="/">Go home</a></p>', nav=site.nav(""), side=site.side("dsa"))

    # rebuild the route folders from scratch so removed pages disappear
    for d in [s["id"] for s in sections] + ["browse"]:
        if d in RESERVED - {"browse"}:
            raise SystemExit(f"section id {d!r} collides with a repo folder")
        shutil.rmtree(root / d, ignore_errors=True)
    for path, h, _, _ in pages:
        f = root / (path.strip("/") if path.endswith(".html") else path.strip("/") + "/index.html").lstrip("/")
        f.parent.mkdir(parents=True, exist_ok=True)
        f.write_text(h, encoding="utf-8", newline="\n")

    # sitemap: <lastmod> moves only when a page's content changes (hashes are kept in state_file)
    state_file = Path(state_file)
    old = json.loads(state_file.read_text(encoding="utf-8")) if state_file.exists() else {}
    today = datetime.now(timezone.utc).strftime("%Y-%m-%d")
    state = {}
    for path, _, digest, noindex in pages:
        if noindex:
            continue
        prev = old.get(path)
        state[path] = prev if prev and prev[0] == digest else [digest, today]
    state_file.write_text(json.dumps(state, indent=0, sort_keys=True) + "\n", encoding="utf-8", newline="\n")
    urls = "".join(f"<url><loc>{ORIGIN}{p}</loc><lastmod>{state[p][1]}</lastmod></url>\n" for p, _, _, ni in pages if not ni)
    (root / "sitemap.xml").write_text('<?xml version="1.0" encoding="UTF-8"?>\n<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">\n'
                                      + urls + "</urlset>\n", encoding="utf-8", newline="\n")
    (root / "robots.txt").write_text(
        "User-agent: *\n"
        "# filtered views (?d=, ?co=, ?tag=…) repeat the unfiltered pages\n"
        "Disallow: /*?\n"
        "# sources and generated runnable files; the pages already show this content\n"
        "Disallow: /content/\nDisallow: /tools/\nDisallow: /python/\nDisallow: /java/\nDisallow: /go/\n\n"
        f"Sitemap: {ORIGIN}/sitemap.xml\n", encoding="utf-8", newline="\n")
    return len(pages)
