/* SDE & MLE Prep – tiny hash-routed SPA.
   Routes: #/  #/browse?filters  #/<section>  #/<section>/<topic>?filters  #/<section>/<topic>/<slug>
   Filters (in the URL so views are shareable): d=Easy,Medium  st=todo|done  tag=heap  topic=graphs  list=grind169 (b75=1 = list=blind75)
   co=Google  win=6m (company window: all time by default)  sort=num|diff|title|freq  q=text */
const REPO = 'https://github.com/mleprep/mleprep.github.io/blob/main/';
const DIFFS = ['Easy', 'Medium', 'Hard'];
const BIG = ['Google', 'Meta', 'Amazon', 'Microsoft', 'Apple', 'Bloomberg'];
const $ = (s, el = document) => el.querySelector(s);
const esc = s => String(s).replace(/[&<>"']/g, c => ({ '&': '&amp;', '<': '&lt;', '>': '&gt;', '"': '&quot;', "'": '&#39;' }[c]));
const store = {
  get(k, d) { try { const v = localStorage.getItem(k); return v == null ? d : JSON.parse(v); } catch (e) { return d; } },
  set(k, v) { try { localStorage.setItem(k, JSON.stringify(v)); } catch (e) { } },
};
const S = {
  m: null, cache: {}, done: store.get('lc-done', {}), lang: store.get('lc-lang', 'python'), tests: store.get('lc-tests', false),
  mode: store.get('lc-mode', 'read'),   // problem pages: 'read' (solution) or 'solve' (editor + in-browser judge)
};
const main = $('#main'), side = $('#side');

const getJSON = url => fetch(url).then(r => { if (!r.ok) throw new Error(url + ' ' + r.status); return r.json(); });
// cache topic files, but forget failed requests so a network blip doesn't break the topic until reload
const topicData = (sec, t) => (S.cache[sec + '/' + t] ||= getJSON(`data/${sec}/${t}.json`).catch(e => { delete S.cache[sec + '/' + t]; throw e; }));
const sec = id => S.m.sections.find(s => s.id === id);
const href = (...p) => '#/' + p.join('/');
const key = (s, it) => s + '/' + it.slug;
const isDone = (s, it) => !!S.done[key(s, it)];
const browse = params => '#/browse' + (params ? '?' + new URLSearchParams(params) : '');

// company tags (data/companies.json, optional): CO.by[slug] = [{c, a, r}] sorted by all-time frequency
let CO = null;
function loadCompanies(d) {
  const by = {};
  for (const [slug, rows] of Object.entries(d.problems)) by[slug] = rows.map(([i, a, r]) => ({ c: d.companies[i], a, r }));
  CO = { by, names: d.companies, source: d.source };
}
const coRows = it => (CO && CO.by[it.slug]) || [];

// study lists (data/lists.json, optional): Blind 75, NeetCode 150/250, Grind 169, LeetCode 75, Top Interview 150
let LISTS = [];
function loadLists(d) {
  LISTS = d.lists.map(l => ({ ...l, pos: new Map(l.slugs.map((s, i) => [s, i])) }));
}
const listById = id => LISTS.find(l => l.id === id);
// membership; without lists.json, Blind 75 still works from the problem tags
const inList = (id, it) => { const l = listById(id); return l ? l.pos.has(it.slug) : id === 'blind75' && (it.tags || []).includes('blind75'); };
// frequency (0-100) of problem `it` at company `c` in window `win` ('' = all time, '6m' = last six months); 0 = not asked
function coVal(it, c, win) {
  const r = coRows(it).find(x => x.c === c);
  return !r ? 0 : win === '6m' ? r.r : (r.a || r.r);
}
// how many companies asked it in the window (used for "Most asked" without a company)
const coCount = (it, win) => coRows(it).filter(x => (win === '6m' ? x.r : (x.a || x.r)) > 0).length;
const fmtF = v => v >= 10 ? Math.round(v) : v.toFixed(1);

// flat index: search, browse, pager, related links
let FLAT = [], BY_ID = {}, TOPIC = {};
function index() {
  FLAT = [];
  S.m.sections.forEach(s => s.topics.forEach((t, ti) => {
    TOPIC[s.id + '/' + t.id] = t.title;
    t.items.forEach((it, ii) => FLAT.push({ s: s.id, t: t.id, tTitle: t.title, it, order: FLAT.length }));
  }));
  for (const f of FLAT) if (f.s === 'dsa') BY_ID[f.it.id] = f;
}

function md(text) {
  const el = document.createElement('div');
  el.className = 'md';
  el.innerHTML = marked.parse(text || '');
  el.querySelectorAll('pre code').forEach(c => {
    if (c.classList.contains('language-viz')) c.parentElement.classList.add('viz');
    else if (window.hljs) hljs.highlightElement(c);
  });
  return el;
}
const hl = (code, lang) => window.hljs ? hljs.highlight(code, { language: lang }).value : esc(code);
const tagLabel = t => t === 'blind75' ? '★ Blind 75' : t === 'premium' ? '🔒 Premium' : t;
const tagClass = t => t === 'blind75' ? ' star' : t === 'premium' ? ' prem' : '';

// ---------------------------------------------------------------- filters
function parseHash() {
  const h = location.hash.replace(/^#\/?/, '');
  const i = h.indexOf('?');
  const path = (i < 0 ? h : h.slice(0, i)).split('/').filter(Boolean).map(decodeURIComponent);
  return { path, params: new URLSearchParams(i < 0 ? '' : h.slice(i + 1)) };
}
function readF(p) {
  return {
    d: new Set((p.get('d') || '').split(',').filter(x => DIFFS.includes(x))),
    st: ['todo', 'done'].includes(p.get('st')) ? p.get('st') : 'all',
    tag: p.get('tag') || '', topic: p.get('topic') || '', list: p.get('list') || (p.get('b75') === '1' ? 'blind75' : ''),
    co: CO ? p.get('co') || '' : '', win: p.get('win') === '6m' ? '6m' : '',
    sort: p.get('sort') || '', q: p.get('q') || '',
  };
}
function writeF(F) {
  const p = new URLSearchParams();
  if (F.q) p.set('q', F.q);
  if (F.d.size) p.set('d', DIFFS.filter(d => F.d.has(d)).join(','));
  if (F.st !== 'all') p.set('st', F.st);
  if (F.list) p.set('list', F.list);
  if (F.tag) p.set('tag', F.tag);
  if (F.topic) p.set('topic', F.topic);
  if (F.co) p.set('co', F.co);
  if (F.win) p.set('win', F.win);
  if (F.sort) p.set('sort', F.sort);
  return p.toString();
}
// does item f pass every filter except `skip` (used for faceted counts)?
function passes(f, F, skip) {
  const it = f.it, tags = it.tags || [];
  if (skip !== 'd' && F.d.size && !F.d.has(it.difficulty)) return false;
  if (skip !== 'st' && F.st !== 'all' && (F.st === 'done') !== isDone(f.s, it)) return false;
  if (skip !== 'list' && F.list && !inList(F.list, it)) return false;
  if (skip !== 'tag' && F.tag && !tags.includes(F.tag)) return false;
  if (skip !== 'topic' && F.topic && f.t !== F.topic) return false;
  if (skip !== 'co' && F.co && !coVal(it, F.co, F.win)) return false;
  if (F.q) {
    const hay = `${it.id || ''} ${it.title} ${tags.join(' ')} ${f.tTitle} ${it.difficulty || ''}`.toLowerCase();
    if (!F.q.toLowerCase().split(/\s+/).filter(Boolean).every(w => hay.includes(w))) return false;
  }
  return true;
}
const DRANK = { Easy: 0, Medium: 1, Hard: 2 };
function sortItems(list, how, F = {}) {
  const a = list.slice();
  if (how === 'freq') {
    const score = F.co ? f => coVal(f.it, F.co, F.win) : f => coCount(f.it, F.win);
    a.sort((x, y) => score(y) - score(x) || x.order - y.order);
  }
  if (how === 'list' && listById(F.list)) {   // the list's own study order (e.g. Grind's week-by-week plan)
    const pos = listById(F.list).pos;
    a.sort((x, y) => (pos.get(x.it.slug) ?? 1e9) - (pos.get(y.it.slug) ?? 1e9));
  }
  if (how === 'num') a.sort((x, y) => (x.it.id || 0) - (y.it.id || 0));
  else if (how === 'diff') a.sort((x, y) => (DRANK[x.it.difficulty] ?? 9) - (DRANK[y.it.difficulty] ?? 9) || x.order - y.order);
  else if (how === 'title') a.sort((x, y) => x.it.title.localeCompare(y.it.title));
  return a;
}

/* Generic filterable list. scope = FLAT entries; base = route path used when filters change. */
function listView(container, scope, base, opts = {}) {
  let F = readF(parseHash().params);
  const hasDiff = scope.some(f => f.it.difficulty), hasTags = scope.some(f => (f.it.tags || []).length);
  const draw = () => {
    const shown = sortItems(scope.filter(f => passes(f, F)), F.sort, F);
    const cnt = (skip, pred) => scope.filter(f => passes(f, F, skip) && pred(f)).length;
    const chip = (attr, val, label, on, n) =>
      `<button class="chip${on ? ' on' : ''}" data-${attr}="${esc(val)}">${label}${n == null ? '' : ` <span>${n}</span>`}</button>`;
    const tagCounts = {};
    scope.filter(f => passes(f, F, 'tag')).forEach(f => (f.it.tags || []).forEach(t => { if (t !== 'blind75') tagCounts[t] = (tagCounts[t] || 0) + 1; }));
    const tagOpts = Object.keys(tagCounts).sort((x, y) => tagCounts[y] - tagCounts[x] || x.localeCompare(y));   // most used first
    if (F.tag && !tagOpts.includes(F.tag)) tagOpts.unshift(F.tag);
    const topics = opts.topicFacet ? S.m.sections[0].topics : [];
    const lists = LISTS.length ? LISTS : scope.some(f => (f.it.tags || []).includes('blind75')) ? [{ id: 'blind75', name: 'Blind 75' }] : [];
    const active = F.d.size || F.st !== 'all' || F.list || F.tag || F.topic || F.co || F.win || F.q;
    // company facet: counts per company among items passing every other filter, in the chosen window
    const hasCo = CO && hasDiff, coCounts = {};
    if (hasCo) scope.filter(f => passes(f, F, 'co')).forEach(f => coRows(f.it).forEach(x => {
      if ((F.win === '6m' ? x.r : (x.a || x.r)) > 0) coCounts[x.c] = (coCounts[x.c] || 0) + 1;
    }));
    const coOpts = Object.keys(coCounts).sort((x, y) => x.localeCompare(y, undefined, { sensitivity: 'base' }));
    if (F.co && !coOpts.includes(F.co)) coOpts.unshift(F.co);
    const winCnt = w => scope.filter(f => passes(f, { ...F, win: w })).length;
    // extra column: the selected company's frequency, or how many companies ask it
    const fcol = hasCo && (F.co || F.sort === 'freq') ? (F.co
      ? f => { const v = coVal(f.it, F.co, F.win); return { w: v, txt: fmtF(v), title: `${F.co} frequency ${fmtF(v)} / 100${F.win ? ' (last 6 months)' : ''}` }; }
      : f => { const n = coCount(f.it, F.win); return { w: 100 * n / (fcolMax || 1), txt: n, title: `asked at ${n} companies${F.win ? ' in the last 6 months' : ''}` }; }) : null;
    const fcolMax = fcol && !F.co ? Math.max(...shown.map(f => coCount(f.it, F.win))) : 0;
    container.innerHTML = `
      <div class="fbar">
        ${hasDiff ? `<div class="fgroup"><span class="flabel">Level</span>${chip('d', '', 'All', !F.d.size)}${DIFFS.map(d =>
          chip('d', d, `<i class="dot ${d}"></i>${d}`, F.d.has(d), cnt('d', f => f.it.difficulty === d))).join('')}</div>` : ''}
        <div class="fgroup"><span class="flabel">Status</span>${chip('st', 'all', 'All', F.st === 'all')}${chip('st', 'todo', 'Todo', F.st === 'todo', cnt('st', f => !isDone(f.s, f.it)))}${chip('st', 'done', 'Done ✓', F.st === 'done', cnt('st', f => isDone(f.s, f.it)))}</div>
        ${hasDiff && lists.length ? `<div class="fgroup"><span class="flabel">List</span>${chip('list', '', 'All', !F.list)}${lists.map(l =>
          chip('list', l.id, l.id === 'blind75' ? '★ Blind 75' : esc(l.name), F.list === l.id, cnt('list', f => inList(l.id, f.it)))).join('')}</div>` : ''}
        ${hasTags ? `<select class="fsel" data-f="tag" aria-label="Pattern tag"><option value="">All patterns</option>${tagOpts.map(t =>
          `<option value="${esc(t)}" ${t === F.tag ? 'selected' : ''}>${esc(t)} (${tagCounts[t] || 0})</option>`).join('')}</select>` : ''}
        ${topics.length ? `<select class="fsel" data-f="topic" aria-label="Topic"><option value="">All topics</option>${topics.map(t =>
          `<option value="${t.id}" ${t.id === F.topic ? 'selected' : ''}>${esc(t.title)}</option>`).join('')}</select>` : ''}
        ${hasDiff ? `<select class="fsel" data-f="sort" aria-label="Sort"><option value="">Roadmap order</option>${[...(listById(F.list) ? [['list', `${listById(F.list).name} order`]] : []), ['num', 'Problem #'], ['diff', 'Difficulty'], ['title', 'Title A–Z'], ...(CO ? [['freq', F.co ? `${F.co} frequency` : 'Most asked']] : [])].map(([v, l]) =>
          `<option value="${v}" ${v === F.sort ? 'selected' : ''}>Sort: ${esc(l)}</option>`).join('')}</select>` : ''}
      </div>
      ${hasCo ? `<div class="fbar">
        <div class="fgroup"><span class="flabel">Company</span>${BIG.filter(c => coCounts[c] || F.co === c).map(c =>
          chip('co', c, esc(c), F.co === c, coCounts[c] || 0)).join('')}
          <select class="fsel" data-f="co" aria-label="Company"><option value="">All companies (${coOpts.length})</option>${coOpts.map(c =>
            `<option value="${esc(c)}" ${c === F.co ? 'selected' : ''}>${esc(c)} (${coCounts[c] || 0})</option>`).join('')}</select></div>
        <div class="fgroup"><span class="flabel">Asked</span>${chip('win', '', 'All time', !F.win, F.co ? winCnt('') : null)}${chip('win', '6m', 'Last 6 months', F.win === '6m', F.co ? winCnt('6m') : null)}</div>
      </div>` : ''}
      <div class="fsum">Showing <b>${shown.length}</b> of ${scope.length}${F.q ? ` for “${esc(F.q)}”` : ''}${active ? ` · <a href="#" data-clear>Clear filters</a>` : ''}</div>
      ${shown.length ? rows(shown, { ...opts, fcol }) : '<p class="muted empty">Nothing matches these filters.</p>'}
      ${hasCo && (F.co || F.sort === 'freq') ? `<p class="muted src">${coSource()}</p>` : ''}`;
    bindChecks(container, draw);
    const update = () => {
      history.replaceState(null, '', '#/' + base + (writeF(F) ? '?' + writeF(F) : ''));
      draw();
    };
    container.querySelectorAll('[data-d]').forEach(b => b.onclick = e => {
      e.preventDefault();
      const d = b.dataset.d;
      if (!d) F.d.clear();
      else if (b.closest('.fbar')) F.d.has(d) ? F.d.delete(d) : F.d.add(d);   // chips toggle (multi-select)
      else { F.d = new Set(F.d.size === 1 && F.d.has(d) ? [] : [d]); }       // clicking a row badge isolates it
      update();
    });
    container.querySelectorAll('[data-st]').forEach(b => b.onclick = () => { F.st = b.dataset.st; update(); });
    container.querySelectorAll('[data-list]').forEach(b => b.onclick = () => {
      F.list = b.dataset.list;
      if (!F.list && F.sort === 'list') F.sort = '';
      update();
    });
    container.querySelectorAll('[data-tag]').forEach(b => b.onclick = e => {
      e.preventDefault();
      if (b.dataset.tag === 'blind75') F.list = 'blind75'; else F.tag = F.tag === b.dataset.tag ? '' : b.dataset.tag;
      update();
    });
    container.querySelectorAll('[data-topic]').forEach(b => b.onclick = e => { e.preventDefault(); F.topic = b.dataset.topic; update(); });
    container.querySelectorAll('[data-co]').forEach(b => b.onclick = e => {
      e.preventDefault();
      F.co = F.co === b.dataset.co ? '' : b.dataset.co;
      if (F.co && !F.sort) F.sort = 'freq';   // picking a company ranks by how often it asks
      update();
    });
    container.querySelectorAll('[data-win]').forEach(b => b.onclick = () => { F.win = b.dataset.win; update(); });
    container.querySelectorAll('select[data-f]').forEach(sel => sel.onchange = () => {
      F[sel.dataset.f] = sel.value;
      if (sel.dataset.f === 'co' && F.co && !F.sort) F.sort = 'freq';
      update();
    });
    const clr = container.querySelector('[data-clear]');
    if (clr) clr.onclick = e => { e.preventDefault(); F = readF(new URLSearchParams()); $('#q').value = ''; update(); };
  };
  draw();
}

function rows(list, opts) {
  const dsa = list.some(f => f.s === 'dsa');
  return `<div class="list-wrap"><table class="list"><tbody>${list.map(f => {
    const it = f.it;
    return `<tr>
      <td class="c"><input type="checkbox" data-k="${key(f.s, it)}" ${isDone(f.s, it) ? 'checked' : ''} aria-label="done"></td>
      ${dsa ? `<td class="n">${it.id || ''}</td>` : ''}
      <td class="t"><a href="${href(f.s, f.t, it.slug)}">${esc(it.title)}</a>
        <span class="tags">${(it.tags || []).map(t => `<a href="${browse(t === 'blind75' ? { list: 'blind75' } : { tag: t })}" class="tag${tagClass(t)}" data-tag="${esc(t)}">${esc(tagLabel(t))}</a>`).join('')}</span></td>
      ${opts.showTopic ? `<td class="tp"><a href="${href(f.s, f.t)}"${f.s === 'dsa' ? ` data-topic="${f.t}"` : ''}>${esc(f.tTitle)}</a></td>` : ''}
      ${opts.fcol ? (v => `<td class="fq" title="${esc(v.title)}"><span class="meter"><i style="width:${Math.max(2, Math.min(100, v.w))}%"></i></span><b>${v.txt}</b></td>`)(opts.fcol(f)) : ''}
      ${it.difficulty ? `<td class="d"><a href="${browse({ d: it.difficulty })}" class="${it.difficulty}" data-d="${it.difficulty}">${it.difficulty}</a></td>` : (dsa ? '<td></td>' : '')}
    </tr>`;
  }).join('')}</tbody></table></div>`;
}
const coSource = () => CO ? `Company data: <a href="${CO.source.repo}" target="_blank" rel="noopener">${esc(CO.source.repo.replace('https://github.com/', ''))}</a>${CO.source.date ? ` (${CO.source.date})` : ''}.
  Frequency is relative within each company (100 = that company's most-asked problem).` : '';
function bindChecks(root, after) {
  root.querySelectorAll('input[data-k]').forEach(cb => cb.onchange = () => {
    if (cb.checked) S.done[cb.dataset.k] = 1; else delete S.done[cb.dataset.k];
    store.set('lc-done', S.done);
    const { path } = parseHash();
    renderSide(path[0] === 'browse' ? 'dsa' : (path[0] || 'dsa'), path[1], path[0] === 'browse');
    if (after) after();
  });
}

// ---------------------------------------------------------------- chrome
function renderNav(cur) {
  $('#sections').innerHTML = S.m.sections.map(s =>
    `<a href="${href(s.id)}" class="${s.id === cur ? 'on' : ''}">${esc(s.title)}</a>`).join('');
}
function renderSide(secId, topicId, onBrowse) {
  const s = sec(secId || 'dsa');
  const total = s.topics.reduce((a, t) => a + t.items.length, 0);
  const done = FLAT.filter(f => f.s === s.id && isDone(s.id, f.it)).length;
  side.innerHTML = `<h4>${esc(s.title)}</h4>` +
    (s.id === 'dsa' ? `<a href="#/browse" class="all${onBrowse ? ' on' : ''}"><span>All problems</span><small>${done ? done + '/' : ''}${total}</small></a>` : '') +
    s.topics.map(t => {
      const d = t.items.filter(it => isDone(s.id, it)).length;
      return `<a href="${href(s.id, t.id)}" title="${esc(t.title)}" class="${t.id === topicId ? 'on' : ''}"><span>${esc(t.title)}</span><small>${d ? d + '/' : ''}${t.items.length}</small></a>`;
    }).join('');
}

// ---------------------------------------------------------------- views
function viewHome() {
  const c = S.m.counts, dsa = sec('dsa');
  const total = dsa.topics.reduce((a, t) => a + t.items.length, 0);
  const done = FLAT.filter(f => f.s === 'dsa' && isDone('dsa', f.it)).length;
  const b75 = FLAT.filter(f => (f.it.tags || []).includes('blind75')).length;
  main.innerHTML = `
    <h1>Interview prep, condensed.</h1>
    <p class="lead">${total} problems — every one from Blind 75, NeetCode 150 &amp; 250, Grind 169, LeetCode 75 and Top Interview 150, plus
    frequently asked company problems — each with the key idea, a sketch, complexity and tested ${c.go >= total ? 'Python, Java &amp; Go'
      : `Python &amp; Java${c.go ? ` (and Go for ${c.go} so far)` : ''}`}. Open any problem and
    pick <b>Solve in browser</b> to write Python${c.go ? ' or Go' : ''} and run it against the tests right on this page (nothing runs on a server).
    Plus system design and ML / deep learning, so one site covers both SDE and MLE interview loops. Progress is saved in this browser.</p>
    <div class="meta big">
      ${DIFFS.map(d => `<a class="badge ${d}" href="${browse({ d })}">${c[d]} ${d}</a>`).join('')}
      <a class="badge star" href="${browse({ list: 'blind75' })}">★ ${b75} Blind 75</a>
      <a href="${browse({ st: 'done' })}">${done}/${total} done</a>
      <a href="${browse({ st: 'todo' })}">Todo list →</a>
      <a href="#" id="rand">Random unsolved →</a>
    </div>
    <div class="cards">
      <a class="card" href="#/browse"><b>Browse all problems</b><span>filter by difficulty, list, pattern, company, status</span></a>
      ${S.m.sections.map(s => {
        const n = s.topics.reduce((a, t) => a + t.items.length, 0);
        return `<a class="card" href="${href(s.id)}"><b>${esc(s.title)}</b><span>${s.topics.length} topics · ${n} ${s.id === 'dsa' ? 'problems' : 'articles'}</span></a>`;
      }).join('')}</div>
    ${LISTS.length ? `<h2>Study lists</h2><div class="cards">${LISTS.map(l => {
      const items = FLAT.filter(f => f.s === 'dsa' && l.pos.has(f.it.slug)), d = items.filter(f => isDone('dsa', f.it)).length;
      return `<a class="card" href="${browse({ list: l.id, sort: 'list' })}"><b>${l.id === 'blind75' ? '★ ' : ''}${esc(l.name)}</b>
        <span>${d}/${items.length} done · in the list's own order</span><div class="bar-prog"><i style="width:${items.length ? 100 * d / items.length : 0}%"></i></div></a>`;
    }).join('')}</div>` : ''}
    ${CO ? `<h2>By company</h2><div class="cards">${BIG.map(c => {
      const all = FLAT.filter(f => f.s === 'dsa' && coVal(f.it, c, '')), recent = all.filter(f => coVal(f.it, c, '6m'));
      return `<a class="card" href="${browse({ co: c, sort: 'freq' })}"><b>${esc(c)}</b><span>${all.length} problems · ${recent.length} in the last 6 months</span></a>`;
    }).join('')}<a class="card" href="${browse({ sort: 'freq' })}"><b>Most asked overall</b><span>ranked by how many companies ask each problem</span></a></div>` : ''}
    <h2>Algorithm roadmap</h2>${topicCards(dsa)}`;
  $('#rand').onclick = e => {
    e.preventDefault();
    const all = FLAT.filter(f => f.s === 'dsa'), todo = all.filter(f => !isDone('dsa', f.it)), pool = todo.length ? todo : all;
    const f = pool[Math.floor(Math.random() * pool.length)];
    location.hash = href(f.s, f.t, f.it.slug);
  };
}
function topicCards(s) {
  return `<div class="cards">${s.topics.map((t, i) => {
    const d = t.items.filter(it => isDone(s.id, it)).length;
    const by = DIFFS.map(x => [x, t.items.filter(it => it.difficulty === x).length]).filter(([, n]) => n);
    return `<a class="card" href="${href(s.id, t.id)}"><b>${s.id === 'dsa' ? (i + 1) + '. ' : ''}${esc(t.title)}</b>
      <span>${d}/${t.items.length} done${by.length ? ' · ' + by.map(([x, n]) => `<em class="${x}">${n}</em>`).join(' ') : ''}</span>
      <div class="bar-prog"><i style="width:${t.items.length ? 100 * d / t.items.length : 0}%"></i></div></a>`;
  }).join('')}</div>`;
}
function viewSection(s) {
  main.innerHTML = `<h1>${esc(s.title)}</h1>${s.id === 'dsa' ? '<p class="lead">Work through the topics in order, or <a href="#/browse">browse and filter all problems</a>.</p>' : ''}${topicCards(s)}`;
}
function viewBrowse() {
  const F = readF(parseHash().params);
  main.innerHTML = `<div class="crumbs"><a href="#/dsa">Algorithms</a></div><h1>${F.q ? 'Search' : 'All problems'}</h1><div id="lv"></div><div id="arts"></div>`;
  listView($('#lv'), FLAT.filter(f => f.s === 'dsa'), 'browse', { topicFacet: true, showTopic: true });
  if (F.q) {   // search also covers system design / ML articles
    const arts = FLAT.filter(f => f.s !== 'dsa' && passes(f, { ...readF(new URLSearchParams()), q: F.q }));
    if (arts.length) $('#arts').innerHTML = `<h2>Articles</h2>${rows(arts, { showTopic: true })}`;
    bindChecks($('#arts'));
  }
}
async function viewTopic(s, t, stale) {
  const data = await topicData(s.id, t.id);
  if (stale()) return;
  main.innerHTML = `<div class="crumbs"><a href="${href(s.id)}">${esc(s.title)}</a></div><h1>${esc(t.title)}</h1>
    <div id="intro" class="article"></div><h2>${s.id === 'dsa' ? 'Problems' : 'Articles'}</h2><div id="lv"></div>`;
  if (data.intro) $('#intro').append(md(data.intro));
  listView($('#lv'), FLAT.filter(f => f.s === s.id && f.t === t.id), `${s.id}/${t.id}`, {});
}
async function viewItem(s, t, slug, stale) {
  const data = await topicData(s.id, t.id);
  if (stale()) return;
  const it = data.items.find(x => x.slug === slug);
  if (!it) return notFound();
  const pos = FLAT.findIndex(f => f.s === s.id && f.t === t.id && f.it.slug === slug);
  const prev = FLAT[pos - 1]?.s === s.id ? FLAT[pos - 1] : null, next = FLAT[pos + 1]?.s === s.id ? FLAT[pos + 1] : null;
  const k = key(s.id, it), dsa = s.id === 'dsa';
  const related = dsa ? it.related.map(id => BY_ID[id]).filter(Boolean) : [];
  const relHtml = related.length ? `<div class="related"><h3>Related problems</h3><div class="chips">${related.map(f =>
    `<a href="${href(f.s, f.t, f.it.slug)}"><i class="dot ${f.it.difficulty}"></i>${f.it.id}. ${esc(f.it.title)}</a>`).join('')}</div></div>` : '';
  main.innerHTML = `
    <div class="crumbs"><a href="${href(s.id)}">${esc(s.title)}</a> › <a href="${href(s.id, t.id)}">${esc(t.title)}</a></div>
    <h1>${dsa ? it.id + '. ' : ''}${esc(it.title)}</h1>
    <div class="meta">
      ${dsa ? `<a class="badge ${it.difficulty}" href="${browse({ d: it.difficulty })}" title="All ${it.difficulty} problems">${it.difficulty}</a>
        <a href="https://leetcode.com/problems/${it.slug}/" target="_blank" rel="noopener">LeetCode ↗</a>` : ''}
      ${(it.tags || []).map(x => `<a class="tag${tagClass(x)}" href="${browse(x === 'blind75' ? { list: 'blind75' } : { tag: x })}" title="Browse problems tagged ${esc(x)}">${esc(tagLabel(x))}</a>`).join('')}
      <label class="done"><input type="checkbox" data-k="${k}" ${S.done[k] ? 'checked' : ''}> Done</label>
    </div>
    ${dsa ? askedAt(it) : ''}
    ${dsa ? `<div class="split">
      <div class="desc"><div id="notes-off" class="notes-off" hidden><b>Solution notes are hidden while you solve.</b>
          Read the problem on <a href="https://leetcode.com/problems/${it.slug}/" target="_blank" rel="noopener">LeetCode ↗</a>${(it.tags || []).includes('premium') ? ' (Premium)' : ''},
          write your solution, and run it against this site's tests. <button id="shownotes" class="chip">Show notes</button></div>
        <div id="body"></div>${relHtml}</div>
      <div class="codecol">
        <div class="modes" role="tablist" aria-label="Mode"><button role="tab" data-mode="read">Solution</button><button role="tab" data-mode="solve">Solve in browser ▶</button></div>
        <div class="code" id="readbox">
          <div class="bar"><button data-l="python">Python</button><button data-l="java">Java</button>${it.go ? '<button data-l="go">Go</button>' : ''}<span class="sp"></span>
            <button id="tests">Tests</button><button id="copy">Copy</button><button id="file" title="Open the runnable file on GitHub">Run locally ↗</button></div>
          <div class="scroll"><pre><code id="src"></code></pre><pre class="test" hidden><code id="tsrc"></code></pre></div></div>
        <div class="code" id="solvebox" hidden>
          <div class="bar"><button data-sl="python" title="Python 3 via Pyodide (CPython compiled to WebAssembly)">Python</button>${it.go ? '<button data-sl="go" title="Go via Yaegi, a Go interpreter compiled to WebAssembly: Go 1.22 semantics, with min/max and slices.Sort available. LeetCode style: no package clause or imports needed.">Go</button>' : ''}<span class="lbl wide" title="Your code runs on your machine, nothing is sent to a server. Java can't be compiled in the browser, so it isn't judged here.">runs in your browser</span><span class="sp"></span>
            <button id="viewtests">Tests</button><button id="reset" title="Replace your code with the starter stub (Ctrl+Z undoes)">Reset</button><button id="run" class="go" title="Ctrl+Enter">Run tests ▶</button></div>
          <pre class="test" id="jtests" hidden><code></code></pre>
          <div id="editor" class="editor"></div></div>
        <div id="results" class="results" aria-live="polite" hidden></div>
      </div>
    </div>` : `<div class="artgrid"><div id="body" class="article"></div>
      <nav class="rail"><h4>In ${esc(t.title)}</h4>${t.items.map(x => `<a href="${href(s.id, t.id, x.slug)}" class="${x.slug === slug ? 'on' : ''}">${esc(x.title)}</a>`).join('')}</nav></div>`}
    <div class="pager"><span>${prev ? `<a href="${href(prev.s, prev.t, prev.it.slug)}">← ${esc(prev.it.title)}</a>` : ''}</span>
      <span>${next ? `<a href="${href(next.s, next.t, next.it.slug)}">${esc(next.it.title)} →</a>` : ''}</span></div>`;
  $('#body').append(md(it.md));
  bindChecks(main);
  const more = $('#comore');
  if (more) more.onclick = () => { $('.asked').classList.add('all'); more.remove(); };
  if (!dsa) {
    if (s.id === 'ml' && t.id === 'ml-coding') addRunButtons($('#body'));
    return;
  }
  const pad = String(it.id).padStart(4, '0');
  const base = pad + '_' + it.slug.replace(/-/g, '_');
  const cls = 'P' + pad + it.slug.split('-').map(w => w[0].toUpperCase() + w.slice(1)).join('');
  // solution + tests + runnable file per language (problems without Go fall back to Python)
  const SRC = { python: [it.py, it.pyTest, `python/${base}.py`], java: [it.java, it.javaTest, `java/${cls}.java`] };
  if (it.go) SRC.go = [it.go, it.goTest, `go/${base}.go`];
  const lang = () => SRC[S.lang] ? S.lang : 'python';
  const show = () => {
    const [src, test] = SRC[lang()];
    $('#src').innerHTML = hl(src, lang());
    $('#tsrc').innerHTML = hl(test || '', lang());
    $('.code pre.test').hidden = !S.tests;
    $('#tests').classList.toggle('on', S.tests);
    main.querySelectorAll('[data-l]').forEach(b => b.classList.toggle('on', b.dataset.l === lang()));
  };
  main.querySelectorAll('[data-l]').forEach(b => b.onclick = () => { S.lang = b.dataset.l; store.set('lc-lang', S.lang); show(); });
  $('#tests').onclick = () => { S.tests = !S.tests; store.set('lc-tests', S.tests); show(); };
  $('#copy').onclick = async () => {
    try { await navigator.clipboard.writeText(SRC[lang()][0]); $('#copy').textContent = 'Copied ✓'; }
    catch (e) { $('#copy').textContent = 'Copy failed'; }
    setTimeout(() => $('#copy') && ($('#copy').textContent = 'Copy'), 1400);
  };
  $('#file').onclick = () => window.open(REPO + SRC[lang()][2], '_blank', 'noopener');
  show();
  solveMode(it, k, t, show);
}

// ---------------------------------------------------------------- solve mode (in-browser judge)
function solveMode(it, k, t, showRead) {
  const L = {
    python: { stub: it.pyStub, test: it.pyTest, draft: 'lc-draft:' + it.slug, judge: PyJudge,
      boot: 'Starting Python… (the first run downloads the runtime, ~10 MB; later runs are instant)' },
  };
  if (it.go) L.go = { stub: it.goStub, test: it.goTest, draft: 'lc-draft-go:' + it.slug, judge: GoJudge,
    boot: 'Starting Go… (the first run downloads the interpreter, ~3 MB; later runs are instant)' };
  const cur = () => L[S.lang] ? S.lang : 'python';   // Java isn't judged: solve in Python
  let ed = null, edLang = null, notes = false, saveT = null, saveFn = null;
  const code = l => ed && edLang === l ? ed.get() : store.get(L[l].draft, null) ?? L[l].stub;
  const flush = () => { if (saveT) { clearTimeout(saveT); saveT = null; saveFn(); } };
  const save = () => {
    const l = edLang;
    clearTimeout(saveT);
    saveFn = () => { saveT = null; const c = code(l); store.set(L[l].draft, c === L[l].stub ? null : c); };
    saveT = setTimeout(saveFn, 300);
  };
  const setLang = () => {
    const l = cur();
    main.querySelectorAll('[data-sl]').forEach(b => b.classList.toggle('on', b.dataset.sl === l));
    const pre = $('#jtests');
    if (!pre.hidden) pre.firstChild.innerHTML = hl(L[l].test, l);
    if (ed && edLang !== l) { flush(); const v = code(l); edLang = l; ed.swap(v, l); }
  };
  const setMode = m => {
    S.mode = m;
    store.set('lc-mode', m);
    const solve = m === 'solve';
    main.querySelectorAll('[data-mode]').forEach(b => { b.classList.toggle('on', b.dataset.mode === m); b.setAttribute('aria-selected', b.dataset.mode === m); });
    $('#readbox').hidden = solve;
    $('#solvebox').hidden = !solve;
    $('.codecol').classList.toggle('solving', solve);
    $('#results').hidden = !solve || !$('#results').innerHTML;
    $('#body').hidden = solve && !notes;
    $('#notes-off').hidden = !solve || notes;
    setLang();
    if (solve) {
      L[cur()].judge.start();   // warm up the runtime while the user reads/types
      if (!ed && !edLang) {
        edLang = cur();
        makeEditor($('#editor'), code(edLang), edLang, { onChange: save, onRun: run }).then(e => { ed = e; setLang(); });
      }
    }
  };
  main.querySelectorAll('[data-sl]').forEach(b => b.onclick = () => {
    if (b.dataset.sl === cur()) return;
    S.lang = b.dataset.sl;
    store.set('lc-lang', S.lang);
    setLang();
    showRead();
    L[cur()].judge.start();
  });
  const run = async () => {
    const btn = $('#run');
    if (!btn) return;
    const l = cur(), J = L[l].judge;
    if (btn.dataset.busy) { PyJudge.kill('Stopped.'); GoJudge.kill('Stopped.'); return; }
    btn.dataset.busy = '1';
    btn.textContent = 'Stop ■';
    const res = $('#results');
    res.hidden = false;
    res.innerHTML = `<div class="rsum">${J.state === 'ready' ? 'Running…' : L[l].boot}</div>`;
    try {
      const r = await J.run({ kind: 'judge', code: code(l), test: L[l].test });
      if (!res.isConnected) return;
      const pass = r.cases.filter(c => c.ok).length, ok = !r.error && r.cases.length && pass === r.cases.length;
      let note = '';
      if (ok && !S.done[k]) {
        S.done[k] = 1;
        store.set('lc-done', S.done);
        main.querySelectorAll(`input[data-k="${k}"]`).forEach(cb => { cb.checked = true; });
        renderSide('dsa', t.id);
        note = ' · marked as done';
      }
      res.innerHTML = judgeHtml(r, note);
    } catch (e) {
      if (res.isConnected) res.innerHTML = `<div class="rsum bad">✗ ${esc(e.message)}</div>`;
    } finally {
      if (btn.isConnected) { delete btn.dataset.busy; btn.textContent = 'Run tests ▶'; }
    }
  };
  main.querySelectorAll('[data-mode]').forEach(b => b.onclick = () => setMode(b.dataset.mode));
  $('#shownotes').onclick = () => { notes = true; setMode('solve'); };
  $('#run').onclick = run;
  $('#reset').onclick = () => {
    const l = cur();
    if (ed && edLang === l) { ed.set(L[l].stub); ed.focus(); }
    flush();
    store.set(L[l].draft, null);
  };
  $('#viewtests').onclick = () => {
    const pre = $('#jtests');
    pre.hidden = !pre.hidden;
    $('#viewtests').classList.toggle('on', !pre.hidden);
    if (!pre.hidden) pre.firstChild.innerHTML = hl(L[cur()].test, cur());
  };
  setMode(S.mode === 'solve' ? 'solve' : 'read');
}

function judgeHtml(r, note = '') {
  const n = r.cases.length, pass = r.cases.filter(c => c.ok).length, ok = !r.error && n && pass === n;
  const head = ok ? '✓ Accepted' : r.error ? '✗ Error' : r.cases.some(c => c.exc) ? '✗ Runtime error' : '✗ Wrong answer';
  const one = s => s.split('\n')[0].slice(0, 120) + (s.includes('\n') || s.length > 120 ? ' …' : '');
  const full = s => s.includes('\n') || s.length > 120 ? `<pre>${esc(s)}</pre>` : '';
  return `<div class="rsum ${ok ? 'good' : 'bad'}">${head}<span>${pass}/${n} tests passed · ${r.ms} ms${note}</span></div>
    ${r.error ? `<pre class="rerr">${esc(r.error)}</pre>` : ''}
    ${r.cases.map((c, i) => `<details class="case ${c.ok ? 'ok' : 'no'}"${c.ok ? '' : ' open'}>
      <summary><b>${c.ok ? '✓' : '✗'}</b><span class="tn">Test ${i + 1}</span><code>${esc(one(c.src))}</code></summary>
      ${full(c.src)}
      ${c.got != null ? `<div class="kv"><span>got</span><code>${esc(c.got)}</code><span>expected</span><code>${esc(c.expected)}</code></div>` : ''}
      ${c.error ? `<div class="kv"><span>error</span><code>${esc(c.error)}</code></div>` : ''}</details>`).join('')}
    ${r.stdout ? `<div class="rout"><span>stdout</span><pre>${esc(r.stdout)}</pre></div>` : ''}`;
}

// ML coding articles: a Run button above each Python snippet (NumPy is loaded on demand)
function addRunButtons(root) {
  root.querySelectorAll('pre > code.language-python').forEach(c => {
    const pre = c.parentElement, src = c.textContent;
    const bar = document.createElement('div');
    bar.className = 'runbar';
    bar.innerHTML = '<button class="go">Run ▶</button><span class="muted">runs in your browser (Python + NumPy via Pyodide)</span>';
    const out = document.createElement('div');
    out.className = 'results';
    out.hidden = true;
    pre.before(bar);
    pre.after(out);
    const btn = bar.querySelector('button');
    btn.onclick = async () => {
      if (btn.dataset.busy) { PyJudge.kill('Stopped.'); return; }
      btn.dataset.busy = '1';
      btn.textContent = 'Stop ■';
      out.hidden = false;
      out.innerHTML = `<div class="rsum">${PyJudge.state === 'ready' ? 'Running…' : 'Starting Python + NumPy… (first run downloads ~15 MB, then it is cached)'}</div>`;
      try {
        const r = await PyJudge.run({ kind: 'script', code: src }, 30000);
        out.innerHTML = `<div class="rsum ${r.error ? 'bad' : 'good'}">${r.error ? '✗ Error' : '✓ Done'}<span>${r.ms} ms</span></div>
          ${r.error ? `<pre class="rerr">${esc(r.error)}</pre>` : ''}${r.stdout ? `<div class="rout"><span>output</span><pre>${esc(r.stdout)}</pre></div>` : ''}`;
      } catch (e) {
        out.innerHTML = `<div class="rsum bad">✗ ${esc(e.message)}</div>`;
      } finally { delete btn.dataset.busy; btn.textContent = 'Run ▶'; }
    };
  });
}

/* Language runtimes in Web Workers: Python (assets/py-worker.js, Pyodide) and Go (assets/go-worker.js, Yaegi).
   A run that exceeds its time limit terminates the worker (the only way to stop an infinite loop); the next
   run starts a new one. */
function makeJudge(url, name) {
  return {
    w: null, seq: 0, pending: new Map(), state: 'idle',
    start() {
      if (this.w) return;
      this.state = 'loading';
      this.w = new Worker(url);
      this.w.onmessage = ({ data }) => {
        if ('ready' in data) {
          this.state = data.ready ? 'ready' : 'error';
          if (!data.ready) this.kill(`Could not start ${name}: ` + data.error);
          return;
        }
        const p = this.pending.get(data.id);
        if (!p) return;
        if (data.started) { p.timer = setTimeout(() => this.kill(`Time limit exceeded (${p.limit / 1000} s) — infinite loop or too slow?`), p.limit); return; }
        clearTimeout(p.timer);
        this.pending.delete(data.id);
        data.ok ? p.resolve(data.result) : p.reject(new Error(data.error));
        if (data.dead) this.kill(data.error);   // the runtime crashed (e.g. stack overflow): start afresh next time
      };
      this.w.onerror = e => { e.preventDefault(); this.kill(`Could not start ${name} (offline or blocked download?)`); };
    },
    run(msg, limit = 10000) {
      this.start();
      return new Promise((resolve, reject) => {
        const id = ++this.seq;
        this.pending.set(id, { resolve, reject, limit, timer: null });
        this.w.postMessage({ id, ...msg });
      });
    },
    kill(reason) {
      if (this.w) this.w.terminate();
      this.w = null;
      this.state = 'idle';
      for (const p of this.pending.values()) { clearTimeout(p.timer); p.reject(new Error(reason)); }
      this.pending.clear();
    },
  };
}
const PyJudge = makeJudge('assets/py-worker.js', 'Python'), GoJudge = makeJudge('assets/go-worker.js', 'Go');

/* Code editor: CodeMirror 5 from cdnjs, loaded on first use; falls back to a plain textarea. */
const CM = 'https://cdnjs.cloudflare.com/ajax/libs/codemirror/5.65.18/';
let cmLoad = null;
const loadScript = src => new Promise((res, rej) => {
  const el = document.createElement('script');
  el.src = src;
  el.onload = res;
  el.onerror = () => rej(new Error('could not load ' + src));
  document.head.append(el);
});
function loadCM() {
  return cmLoad ||= (async () => {
    const css = document.createElement('link');
    css.rel = 'stylesheet';
    css.href = CM + 'codemirror.min.css';
    document.head.prepend(css);   // before style.css so the site's editor theme wins
    await loadScript(CM + 'codemirror.min.js');
    await Promise.all(['mode/python/python.min.js', 'mode/go/go.min.js', 'addon/edit/matchbrackets.min.js',
      'addon/edit/closebrackets.min.js', 'addon/comment/comment.min.js'].map(p => loadScript(CM + p)));
  })().catch(e => { cmLoad = null; throw e; });
}
// Python indents with 4 spaces, Go (gofmt) with tabs
const EDITOR = { python: { mode: 'python', tabs: false, label: 'Python code editor' }, go: { mode: 'go', tabs: true, label: 'Go code editor' } };
async function makeEditor(host, value, lang, { onChange, onRun }) {
  try {
    await loadCM();
    const cm = CodeMirror(host, {
      value, mode: EDITOR[lang].mode, indentUnit: 4, tabSize: 4, indentWithTabs: EDITOR[lang].tabs, lineNumbers: true,
      matchBrackets: true, autoCloseBrackets: true, viewportMargin: Infinity, screenReaderLabel: EDITOR[lang].label,
      extraKeys: {
        Tab: c => c.somethingSelected() ? c.indentSelection('add') : c.replaceSelection(c.getOption('indentWithTabs') ? '\t' : '    '),
        'Shift-Tab': c => c.indentSelection('subtract'),
        'Ctrl-Enter': () => onRun(), 'Cmd-Enter': () => onRun(),
        'Ctrl-/': 'toggleComment', 'Cmd-/': 'toggleComment',
      },
    });
    cm.on('change', onChange);
    return {
      get: () => cm.getValue(), set: v => cm.setValue(v), focus: () => cm.focus(),
      swap(v, l) {   // switch language: new mode, new text, fresh undo history
        cm.setOption('mode', EDITOR[l].mode);
        cm.setOption('indentWithTabs', EDITOR[l].tabs);
        cm.setOption('screenReaderLabel', EDITOR[l].label);
        cm.setValue(v);
        cm.clearHistory();
      },
    };
  } catch (e) {
    const ta = document.createElement('textarea');
    let tab = EDITOR[lang].tabs ? '\t' : '    ';
    ta.className = 'plain';
    ta.spellcheck = false;
    ta.value = value;
    ta.setAttribute('aria-label', EDITOR[lang].label);
    ta.addEventListener('input', onChange);
    ta.addEventListener('keydown', ev => {
      if (ev.key === 'Enter' && (ev.ctrlKey || ev.metaKey)) { ev.preventDefault(); onRun(); }
      else if (ev.key === 'Tab' && !ev.shiftKey) { ev.preventDefault(); ta.setRangeText(tab, ta.selectionStart, ta.selectionEnd, 'end'); onChange(); }
    });
    host.append(ta);
    return {
      get: () => ta.value, set: v => { ta.value = v; }, focus: () => ta.focus(),
      swap(v, l) { ta.value = v; tab = EDITOR[l].tabs ? '\t' : '    '; ta.setAttribute('aria-label', EDITOR[l].label); },
    };
  }
}
// "Asked at" chips on a problem page: companies by all-time frequency; ● = also asked in the last six months
function askedAt(it) {
  const tier = x => BIG.includes(x.c) ? 0 : x.r ? 1 : 2;   // big tech, then recently asked, then the rest
  const list = coRows(it).slice().sort((x, y) => tier(x) - tier(y) || (y.a || y.r) - (x.a || x.r));
  if (!list.length) return '';
  const TOP = 12;
  return `<div class="asked"><span class="flabel" title="${esc(coSource().replace(/<[^>]+>/g, ''))}">Asked at ${list.length}</span>${list.map((x, i) =>
    `<a class="co${i >= TOP ? ' extra' : ''}${x.r ? ' hot' : ''}" href="${browse({ co: x.c, sort: 'freq' })}"
      title="${esc(x.c)}: frequency ${fmtF(x.a)} all time${x.r ? `, ${fmtF(x.r)} in the last 6 months` : ''}">${esc(x.c)}<span>${fmtF(x.a || x.r)}</span></a>`).join('')}${list.length > TOP ? `<button id="comore" class="co more">+${list.length - TOP} more</button>` : ''}</div>`;
}
function notFound() { main.innerHTML = `<h1>Not found</h1><p><a href="#/">Go home</a></p>`; }

// ---------------------------------------------------------------- router
let lastPath = null, routeSeq = 0;
async function route() {
  const my = ++routeSeq, stale = () => my !== routeSeq;   // a newer navigation started while we awaited data
  document.body.classList.remove('open');
  const { path, params } = parseHash();
  const [sid, tid, slug] = path;
  const onBrowse = sid === 'browse';
  const s = !onBrowse && sid && sec(sid), t = s && tid && s.topics.find(x => x.id === tid);
  renderNav(onBrowse ? 'dsa' : s ? s.id : '');
  renderSide(s ? s.id : 'dsa', t && t.id, onBrowse);
  if (document.activeElement !== $('#q')) $('#q').value = onBrowse ? params.get('q') || '' : '';
  try {
    if (!sid) viewHome();
    else if (onBrowse) viewBrowse();
    else if (!s) notFound();
    else if (!tid) viewSection(s);
    else if (!t) notFound();
    else if (!slug) await viewTopic(s, t, stale);
    else await viewItem(s, t, slug, stale);
  } catch (e) { if (!stale()) main.innerHTML = `<h1>Error</h1><pre>${esc(e.message)}</pre>`; }
  if (stale()) return;
  document.title = (main.querySelector('h1')?.textContent || '') + ' · SDE & MLE Prep';
  const p = path.join('/');
  if (p !== lastPath) window.scrollTo(0, 0);   // keep scroll when only filters change
  lastPath = p;
}

// search box = browse view with ?q=
let qTimer, preSearch = null;
$('#q').addEventListener('input', e => {
  clearTimeout(qTimer);
  qTimer = setTimeout(() => {
    const q = e.target.value.trim();
    if (!location.hash.startsWith('#/browse')) preSearch = location.hash;
    if (q) {
      const F = readF(parseHash().path[0] === 'browse' ? parseHash().params : new URLSearchParams());
      F.q = q;
      history.replaceState(null, '', '#/browse?' + writeF(F));
      route();
    } else if (preSearch) { location.hash = preSearch; preSearch = null; }
    else { const F = readF(parseHash().params); F.q = ''; history.replaceState(null, '', '#/browse' + (writeF(F) ? '?' + writeF(F) : '')); route(); }
  }, 150);
});
document.addEventListener('keydown', e => {
  const a = document.activeElement;
  if (e.key === '/' && a !== $('#q') && !/INPUT|SELECT|TEXTAREA/.test(a.tagName) && !a.isContentEditable) { e.preventDefault(); $('#q').focus(); }
  if (e.key === 'Escape' && document.activeElement === $('#q')) { $('#q').value = ''; $('#q').dispatchEvent(new Event('input')); $('#q').blur(); }
});
$('#menu').onclick = () => document.body.classList.toggle('open');
$('#theme').onclick = () => {
  const dark = matchMedia('(prefers-color-scheme: dark)').matches;
  const cur = document.documentElement.dataset.theme || (dark ? 'dark' : 'light');
  const nxt = cur === 'dark' ? 'light' : 'dark';
  document.documentElement.dataset.theme = nxt;
  store.set('lc-theme', nxt);
};
window.addEventListener('hashchange', route);

Promise.all([getJSON('data/manifest.json'), getJSON('data/companies.json').catch(() => null), getJSON('data/lists.json').catch(() => null)])
  .then(([m, co, lists]) => { S.m = m; if (co) loadCompanies(co); if (lists) loadLists(lists); index(); route(); })
  .catch(e => { main.innerHTML = `<h1>Could not load data</h1><pre>${esc(e.message)}</pre>`; });
