/* LeetCode Review – tiny hash-routed SPA.
   Routes: #/  #/browse?filters  #/<section>  #/<section>/<topic>?filters  #/<section>/<topic>/<slug>
   Filters (in the URL so views are shareable): d=Easy,Medium  st=todo|done  tag=heap  topic=graphs  b75=1  sort=num|diff|title  q=text */
const REPO = 'https://github.com/mleprep/mleprep.github.io/blob/main/';
const DIFFS = ['Easy', 'Medium', 'Hard'];
const $ = (s, el = document) => el.querySelector(s);
const esc = s => String(s).replace(/[&<>"']/g, c => ({ '&': '&amp;', '<': '&lt;', '>': '&gt;', '"': '&quot;', "'": '&#39;' }[c]));
const store = {
  get(k, d) { try { const v = localStorage.getItem(k); return v == null ? d : JSON.parse(v); } catch (e) { return d; } },
  set(k, v) { try { localStorage.setItem(k, JSON.stringify(v)); } catch (e) { } },
};
const S = { m: null, cache: {}, done: store.get('lc-done', {}), lang: store.get('lc-lang', 'python'), tests: store.get('lc-tests', false) };
const main = $('#main'), side = $('#side');

const getJSON = url => fetch(url).then(r => { if (!r.ok) throw new Error(url + ' ' + r.status); return r.json(); });
const topicData = (sec, t) => (S.cache[sec + '/' + t] ||= getJSON(`data/${sec}/${t}.json`));
const sec = id => S.m.sections.find(s => s.id === id);
const href = (...p) => '#/' + p.join('/');
const key = (s, it) => s + '/' + it.slug;
const isDone = (s, it) => !!S.done[key(s, it)];
const browse = params => '#/browse' + (params ? '?' + new URLSearchParams(params) : '');

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
const tagLabel = t => t === 'blind75' ? '★ Blind 75' : t;

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
    tag: p.get('tag') || '', topic: p.get('topic') || '', b75: p.get('b75') === '1',
    sort: p.get('sort') || '', q: p.get('q') || '',
  };
}
function writeF(F) {
  const p = new URLSearchParams();
  if (F.q) p.set('q', F.q);
  if (F.d.size) p.set('d', DIFFS.filter(d => F.d.has(d)).join(','));
  if (F.st !== 'all') p.set('st', F.st);
  if (F.b75) p.set('b75', '1');
  if (F.tag) p.set('tag', F.tag);
  if (F.topic) p.set('topic', F.topic);
  if (F.sort) p.set('sort', F.sort);
  return p.toString();
}
// does item f pass every filter except `skip` (used for faceted counts)?
function passes(f, F, skip) {
  const it = f.it, tags = it.tags || [];
  if (skip !== 'd' && F.d.size && !F.d.has(it.difficulty)) return false;
  if (skip !== 'st' && F.st !== 'all' && (F.st === 'done') !== isDone(f.s, it)) return false;
  if (skip !== 'b75' && F.b75 && !tags.includes('blind75')) return false;
  if (skip !== 'tag' && F.tag && !tags.includes(F.tag)) return false;
  if (skip !== 'topic' && F.topic && f.t !== F.topic) return false;
  if (F.q) {
    const hay = `${it.id || ''} ${it.title} ${tags.join(' ')} ${f.tTitle} ${it.difficulty || ''}`.toLowerCase();
    if (!F.q.toLowerCase().split(/\s+/).filter(Boolean).every(w => hay.includes(w))) return false;
  }
  return true;
}
const DRANK = { Easy: 0, Medium: 1, Hard: 2 };
function sortItems(list, how) {
  const a = list.slice();
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
    const shown = sortItems(scope.filter(f => passes(f, F)), F.sort);
    const cnt = (skip, pred) => scope.filter(f => passes(f, F, skip) && pred(f)).length;
    const chip = (attr, val, label, on, n) =>
      `<button class="chip${on ? ' on' : ''}" data-${attr}="${esc(val)}">${label}${n == null ? '' : ` <span>${n}</span>`}</button>`;
    const tagCounts = {};
    scope.filter(f => passes(f, F, 'tag')).forEach(f => (f.it.tags || []).forEach(t => { if (t !== 'blind75') tagCounts[t] = (tagCounts[t] || 0) + 1; }));
    const tagOpts = Object.keys(tagCounts).sort();
    if (F.tag && !tagOpts.includes(F.tag)) tagOpts.unshift(F.tag);
    const topics = opts.topicFacet ? S.m.sections[0].topics : [];
    const active = F.d.size || F.st !== 'all' || F.b75 || F.tag || F.topic || F.q;
    container.innerHTML = `
      <div class="fbar">
        ${hasDiff ? `<div class="fgroup"><span class="flabel">Level</span>${chip('d', '', 'All', !F.d.size)}${DIFFS.map(d =>
          chip('d', d, `<i class="dot ${d}"></i>${d}`, F.d.has(d), cnt('d', f => f.it.difficulty === d))).join('')}</div>` : ''}
        <div class="fgroup"><span class="flabel">Status</span>${chip('st', 'all', 'All', F.st === 'all')}${chip('st', 'todo', 'Todo', F.st === 'todo', cnt('st', f => !isDone(f.s, f.it)))}${chip('st', 'done', 'Done ✓', F.st === 'done', cnt('st', f => isDone(f.s, f.it)))}</div>
        ${hasTags && scope.some(f => (f.it.tags || []).includes('blind75')) ? `<div class="fgroup">${chip('b75', '1', '★ Blind 75', F.b75, cnt('b75', f => (f.it.tags || []).includes('blind75')))}</div>` : ''}
        ${hasTags ? `<select class="fsel" data-f="tag" aria-label="Pattern tag"><option value="">All patterns</option>${tagOpts.map(t =>
          `<option value="${esc(t)}" ${t === F.tag ? 'selected' : ''}>${esc(t)} (${tagCounts[t] || 0})</option>`).join('')}</select>` : ''}
        ${topics.length ? `<select class="fsel" data-f="topic" aria-label="Topic"><option value="">All topics</option>${topics.map(t =>
          `<option value="${t.id}" ${t.id === F.topic ? 'selected' : ''}>${esc(t.title)}</option>`).join('')}</select>` : ''}
        ${hasDiff ? `<select class="fsel" data-f="sort" aria-label="Sort"><option value="">Roadmap order</option>${[['num', 'Problem #'], ['diff', 'Difficulty'], ['title', 'Title A–Z']].map(([v, l]) =>
          `<option value="${v}" ${v === F.sort ? 'selected' : ''}>Sort: ${l}</option>`).join('')}</select>` : ''}
      </div>
      <div class="fsum">Showing <b>${shown.length}</b> of ${scope.length}${F.q ? ` for “${esc(F.q)}”` : ''}${active ? ` · <a href="#" data-clear>Clear filters</a>` : ''}</div>
      ${shown.length ? rows(shown, opts) : '<p class="muted empty">Nothing matches these filters.</p>'}`;
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
    container.querySelectorAll('[data-b75]').forEach(b => b.onclick = () => { F.b75 = !F.b75; update(); });
    container.querySelectorAll('[data-tag]').forEach(b => b.onclick = e => {
      e.preventDefault();
      if (b.dataset.tag === 'blind75') F.b75 = true; else F.tag = F.tag === b.dataset.tag ? '' : b.dataset.tag;
      update();
    });
    container.querySelectorAll('[data-topic]').forEach(b => b.onclick = e => { e.preventDefault(); F.topic = b.dataset.topic; update(); });
    container.querySelectorAll('select[data-f]').forEach(sel => sel.onchange = () => { F[sel.dataset.f] = sel.value; update(); });
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
        <span class="tags">${(it.tags || []).map(t => `<a href="#" class="tag${t === 'blind75' ? ' star' : ''}" data-tag="${esc(t)}">${esc(tagLabel(t))}</a>`).join('')}</span></td>
      ${opts.showTopic ? `<td class="tp"><a href="#" data-topic="${f.t}">${esc(f.tTitle)}</a></td>` : ''}
      ${it.difficulty ? `<td class="d"><a href="#" class="${it.difficulty}" data-d="${it.difficulty}">${it.difficulty}</a></td>` : (dsa ? '<td></td>' : '')}
    </tr>`;
  }).join('')}</tbody></table></div>`;
}
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
    <p class="lead">${total} hand-picked problems (Blind 75, NeetCode, Grind 169, Top Interview 150) with the key idea, a sketch,
    complexity and runnable Python &amp; Java — plus system design and ML / deep learning. Progress is saved in this browser.</p>
    <div class="meta big">
      ${DIFFS.map(d => `<a class="badge ${d}" href="${browse({ d })}">${c[d]} ${d}</a>`).join('')}
      <a class="badge star" href="${browse({ b75: 1 })}">★ ${b75} Blind 75</a>
      <a href="${browse({ st: 'done' })}">${done}/${total} done</a>
      <a href="${browse({ st: 'todo' })}">Todo list →</a>
      <a href="#" id="rand">Random unsolved →</a>
    </div>
    <div class="cards">
      <a class="card" href="#/browse"><b>Browse all problems</b><span>filter by difficulty, pattern, status</span></a>
      ${S.m.sections.map(s => {
        const n = s.topics.reduce((a, t) => a + t.items.length, 0);
        return `<a class="card" href="${href(s.id)}"><b>${esc(s.title)}</b><span>${s.topics.length} topics · ${n} ${s.id === 'dsa' ? 'problems' : 'articles'}</span></a>`;
      }).join('')}</div>
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
async function viewTopic(s, t) {
  const data = await topicData(s.id, t.id);
  main.innerHTML = `<div class="crumbs"><a href="${href(s.id)}">${esc(s.title)}</a></div><h1>${esc(t.title)}</h1>
    <div id="intro" class="article"></div><h2>${s.id === 'dsa' ? 'Problems' : 'Articles'}</h2><div id="lv"></div>`;
  if (data.intro) $('#intro').append(md(data.intro));
  listView($('#lv'), FLAT.filter(f => f.s === s.id && f.t === t.id), `${s.id}/${t.id}`, {});
}
async function viewItem(s, t, slug) {
  const data = await topicData(s.id, t.id);
  const it = data.items.find(x => x.slug === slug);
  if (!it) return notFound();
  const pos = FLAT.findIndex(f => f.s === s.id && f.it.slug === slug);
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
      ${(it.tags || []).map(x => `<a class="tag${x === 'blind75' ? ' star' : ''}" href="${browse(x === 'blind75' ? { b75: 1 } : { tag: x })}" title="Browse problems tagged ${esc(x)}">${esc(tagLabel(x))}</a>`).join('')}
      <label class="done"><input type="checkbox" data-k="${k}" ${S.done[k] ? 'checked' : ''}> Done</label>
    </div>
    ${dsa ? `<div class="split">
      <div class="desc"><div id="body"></div>${relHtml}</div>
      <div class="codecol"><div class="code">
        <div class="bar"><button data-l="python">Python</button><button data-l="java">Java</button><span class="sp"></span>
          <button id="tests">Tests</button><button id="copy">Copy</button><button id="file" title="Open the runnable file on GitHub">Run locally ↗</button></div>
        <div class="scroll"><pre><code id="src"></code></pre><pre class="test" hidden><code id="tsrc"></code></pre></div></div></div>
    </div>` : `<div class="artgrid"><div id="body" class="article"></div>
      <nav class="rail"><h4>In ${esc(t.title)}</h4>${t.items.map(x => `<a href="${href(s.id, t.id, x.slug)}" class="${x.slug === slug ? 'on' : ''}">${esc(x.title)}</a>`).join('')}</nav></div>`}
    <div class="pager"><span>${prev ? `<a href="${href(prev.s, prev.t, prev.it.slug)}">← ${esc(prev.it.title)}</a>` : ''}</span>
      <span>${next ? `<a href="${href(next.s, next.t, next.it.slug)}">${esc(next.it.title)} →</a>` : ''}</span></div>`;
  $('#body').append(md(it.md));
  bindChecks(main);
  if (!dsa) return;
  const pad = String(it.id).padStart(4, '0');
  const base = pad + '_' + it.slug.replace(/-/g, '_');
  const cls = 'P' + pad + it.slug.split('-').map(w => w[0].toUpperCase() + w.slice(1)).join('');
  const show = () => {
    const py = S.lang === 'python';
    $('#src').innerHTML = hl(py ? it.py : it.java, S.lang);
    $('#tsrc').innerHTML = hl((py ? it.pyTest : it.javaTest) || '', S.lang);
    $('.code pre.test').hidden = !S.tests;
    $('#tests').classList.toggle('on', S.tests);
    main.querySelectorAll('[data-l]').forEach(b => b.classList.toggle('on', b.dataset.l === S.lang));
  };
  main.querySelectorAll('[data-l]').forEach(b => b.onclick = () => { S.lang = b.dataset.l; store.set('lc-lang', S.lang); show(); });
  $('#tests').onclick = () => { S.tests = !S.tests; store.set('lc-tests', S.tests); show(); };
  $('#copy').onclick = async () => {
    try { await navigator.clipboard.writeText(S.lang === 'python' ? it.py : it.java); $('#copy').textContent = 'Copied ✓'; }
    catch (e) { $('#copy').textContent = 'Copy failed'; }
    setTimeout(() => $('#copy') && ($('#copy').textContent = 'Copy'), 1400);
  };
  $('#file').onclick = () => window.open(REPO + (S.lang === 'python' ? `python/${base}.py` : `java/${cls}.java`), '_blank', 'noopener');
  show();
}
function notFound() { main.innerHTML = `<h1>Not found</h1><p><a href="#/">Go home</a></p>`; }

// ---------------------------------------------------------------- router
let lastPath = null;
async function route() {
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
    else if (!slug) await viewTopic(s, t);
    else await viewItem(s, t, slug);
  } catch (e) { main.innerHTML = `<h1>Error</h1><pre>${esc(e.message)}</pre>`; }
  document.title = (main.querySelector('h1')?.textContent || '') + ' · LeetCode Review';
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
  if (e.key === '/' && document.activeElement !== $('#q') && !/INPUT|SELECT|TEXTAREA/.test(document.activeElement.tagName)) { e.preventDefault(); $('#q').focus(); }
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

getJSON('data/manifest.json').then(m => { S.m = m; index(); route(); })
  .catch(e => { main.innerHTML = `<h1>Could not load data</h1><pre>${esc(e.message)}</pre>`; });
