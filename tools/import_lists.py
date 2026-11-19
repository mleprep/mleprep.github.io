#!/usr/bin/env python3
"""Build data/lists.json: membership and order of the well-known study lists, fetched from public sources.

    python tools/import_lists.py

Sources (no login needed):
  Blind 75, NeetCode 150  github.com/neetcode-gh/leetcode  .problemSiteData.json (blind75 / neetcode150 flags)
  Grind 169               techinterviewhandbook.org/grind75 (problem data embedded in the page's JS bundle)
  LeetCode 75, Top Interview 150   leetcode.com study plans (public GraphQL)
Prints which list problems are missing from content/dsa so they can be added.
"""
import json
import re
import sys
import urllib.request
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
UA = {"User-Agent": "Mozilla/5.0 (study-list importer)"}


def get(url, data=None, headers=None):
    req = urllib.request.Request(url, data=data, headers={**UA, **(headers or {})})
    with urllib.request.urlopen(req, timeout=60) as r:
        return r.read().decode("utf-8", "replace")


def neetcode():
    d = json.loads(get("https://raw.githubusercontent.com/neetcode-gh/leetcode/main/.problemSiteData.json"))
    slug = lambda x: x["link"].strip("/").split("/")[-1]
    return [slug(x) for x in d if x.get("blind75")], [slug(x) for x in d if x.get("neetcode150")]


def grind():
    base = "https://www.techinterviewhandbook.org"
    page = get(base + "/grind75/?weeks=26&hours=40&grouping=none")
    for src in re.findall(r'src="(/grind75/_next/[^"]+\.js)"', page):
        js = get(base + src)
        m = re.search(r"JSON\.parse\('(\[\{\"slug\".*?\])'\)", js, re.S)
        if m:
            return [q["slug"] for q in json.loads(m.group(1).encode().decode("unicode_escape"))]
    sys.exit("Grind data not found (site layout changed?)")


def study_plan(slug):
    q = 'query q($s:String!){studyPlanV2Detail(planSlug:$s){planSubGroups{questions{titleSlug}}}}'
    body = json.dumps({"query": q, "variables": {"s": slug}}).encode()
    d = json.loads(get("https://leetcode.com/graphql/", body, {"Content-Type": "application/json", "Referer": "https://leetcode.com"}))
    return [x["titleSlug"] for g in d["data"]["studyPlanV2Detail"]["planSubGroups"] for x in g["questions"]]


def main():
    b75, nc150 = neetcode()
    lists = [
        {"id": "blind75", "name": "Blind 75", "url": "https://www.teamblind.com/post/New-Year-Gift---Curated-List-of-Top-75-LeetCode-Questions-to-Save-Your-Time-OaM1orEU", "slugs": b75},
        {"id": "neetcode150", "name": "NeetCode 150", "url": "https://neetcode.io/practice", "slugs": nc150},
        {"id": "grind169", "name": "Grind 169", "url": "https://www.techinterviewhandbook.org/grind75", "slugs": grind()},
        {"id": "lc75", "name": "LeetCode 75", "url": "https://leetcode.com/studyplan/leetcode-75/", "slugs": study_plan("leetcode-75")},
        {"id": "ti150", "name": "Top Interview 150", "url": "https://leetcode.com/studyplan/top-interview-150/", "slugs": study_plan("top-interview-150")},
    ]
    m = json.loads((ROOT / "data" / "manifest.json").read_text(encoding="utf-8"))
    ours = {it["slug"] for s in m["sections"] if s["id"] == "dsa" for t in s["topics"] for it in t["items"]}
    for l in lists:
        missing = [s for s in l["slugs"] if s not in ours]
        print(f"{l['name']}: {len(l['slugs'])} problems, {len(l['slugs']) - len(missing)} on the site"
              + (f"; missing: {', '.join(missing)}" if missing else ""))
    (ROOT / "data" / "lists.json").write_text(json.dumps({"lists": lists}, separators=(",", ":")), encoding="utf-8")


if __name__ == "__main__":
    main()
