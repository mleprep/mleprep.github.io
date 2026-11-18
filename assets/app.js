/* LeetCode Review – tiny hash-routed SPA. Routes: #/  #/<section>  #/<section>/<topic>  #/<section>/<topic>/<slug> */
const REPO = 'https://github.com/mleprep/mleprep.github.io/blob/main/';
const $ = (s, el = document) => el.querySelector(s);
const esc = s => String(s).replace(/[&<>"']/g, c => ({ '&': '&amp;', '<': '&lt;', '>': '&gt;', '"': '&quot;', "'": '&#39;' }[c]));
const store = {
  get(k, d) { try { const v = localStorage.getItem(k); return v == null ? d : JSON.parse(v); } catch (e) { return d; } },
  set(k, v) { try { localStorage.setItem(k, JSON.stringify(v)); } catch (e) { } },
};
const S = { m: null, cache: {}, done: store.get('lc-done', {}), lang: store.get('lc-lang', 'python'), tests: store.get('lc-tests', false), diff: 'All' };
const main = $('#main'), side = $('#side');

const getJSON = url => fetch(url).then(r => { if (!r.ok) throw new Error(url + ' ' + r.status); return r.json(); });
const topicData = (sec, t) => (S.cache[sec + '/' + t] ||= getJSON(`data/${sec}/${t}.json`));
const sec = id => S.m.sections.find(s => s.id === id);
const href = (...p) => '#/' + p.join('/');
const key = (s, it) => s + '/' + it.slug;
const isDone = (s, it) => !!S.done[key(s, it)];

// flat index for search, pager and related links
let FLAT = [], BY_ID = {};
function index() {
  FLAT = [];
  for (const s of S.m.sections) for (const t of s.topics) for (const it of t.items) FLAT.push({ s: s.id, t: t.id, tTitle: t.title, it });
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
function hl(code, lang) {
  return window.hljs ? hljs.highlight(code, { language: lang }).value : esc(code);
}

// ---------------------------------------------------------------- chrome
function renderNav(cur) {
  $('#sections').innerHTML = S.m.sections.map(s =>
    `<a href="${href(s.id)}" class="${s.id === cur ? 'on' : ''}">${esc(s.title)}</a>`).join('');
}
function renderSide(secId, topicId) {
  const s = sec(secId || 'dsa');
  side.innerHTML = `<h4>${esc(s.title)}</h4>` + s.topics.map(t => {
    const d = t.items.filter(it => isDone(s.id, it)).length;
    return `<a href="${href(s.id, t.id)}" class="${t.id === topicId ? 'on' : ''}"><span>${esc(t.title)}</span><small>${d ? d + '/' : ''}${t.items.length}</small></a>`;
  }).join('');
}

// ---------------------------------------------------------------- views
function viewHome() {
  const c = S.m.counts, dsa = sec('dsa');
  const total = dsa.topics.reduce((a, t) => a + t.items.length, 0);
  const done = FLAT.filter(f => f.s === 'dsa' && isDone('dsa', f.it)).length;
  main.innerHTML = `
    <h1>Interview prep, condensed.</h1>
    <p class="muted">${total} hand-picked problems (Blind 75, NeetCode 150, Grind 169, Top Interview 150) with the key idea, a sketch,
    complexity and runnable Python &amp; Java, plus system design and ML / deep-learning. Progress is saved in this browser.</p>
    <div class="meta"><span class="badge Easy">${c.Easy} Easy</span><span class="badge Medium">${c.Medium} Medium</span><span class="badge Hard">${c.Hard} Hard</span>
    <span>${done}/${total} done</span><a href="#" id="rand">Random problem →</a></div>
    <div class="cards">${S.m.sections.map(s => {
      const n = s.topics.reduce((a, t) => a + t.items.length, 0);
      return `<a class="card" href="${href(s.id)}"><b>${esc(s.title)}</b><span>${s.topics.length} topics · ${n} ${s.id === 'dsa' ? 'problems' : 'articles'}</span></a>`;
    }).join('')}</div>
    <h2>Algorithm roadmap</h2>${topicCards(dsa)}`;
  $('#rand').onclick = e => {
    e.preventDefault();
    const todo = FLAT.filter(f => f.s === 'dsa' && !isDone('dsa', f.it)), pool = todo.length ? todo : FLAT.filter(f => f.s === 'dsa');
    const f = pool[Math.floor(Math.random() * pool.length)];
    location.hash = href(f.s, f.t, f.it.slug);
  };
}
function topicCards(s) {
  return `<div class="cards">${s.topics.map((t, i) => {
    const d = t.items.filter(it => isDone(s.id, it)).length;
    return `<a class="card" href="${href(s.id, t.id)}"><b>${s.id === 'dsa' ? (i + 1) + '. ' : ''}${esc(t.title)}</b><span>${d}/${t.items.length} done</span>
      <div class="bar-prog"><i style="width:${t.items.length ? 100 * d / t.items.length : 0}%"></i></div></a>`;
  }).join('')}</div>`;
}
function viewSection(s) {
  main.innerHTML = `<h1>${esc(s.title)}</h1>${topicCards(s)}`;
}
function rows(s, t, items) {
  return `<table class="list">${items.map(it => `<tr>
    <td class="c"><input type="checkbox" data-k="${key(s.id, it)}" ${isDone(s.id, it) ? 'checked' : ''} aria-label="done"></td>
    ${s.id === 'dsa' ? `<td class="n">${it.id}</td>` : ''}
    <td><a href="${href(s.id, t, it.slug)}">${esc(it.title)}</a><span class="tags">${tags(it.tags)}</span></td>
    ${it.difficulty ? `<td class="d ${it.difficulty}">${it.difficulty}</td>` : ''}</tr>`).join('')}</table>`;
}
function tags(list) {
  return (list || []).map(x => x === 'blind75' ? '<span class="tag star">★ Blind 75</span>' : `<span class="tag">${esc(x)}</span>`).join('');
}
function bindChecks(root) {
  root.querySelectorAll('input[data-k]').forEach(cb => cb.onchange = () => {
    if (cb.checked) S.done[cb.dataset.k] = 1; else delete S.done[cb.dataset.k];
    store.set('lc-done', S.done);
    const [s] = location.hash.slice(2).split('/');
    renderSide(s || 'dsa', location.hash.slice(2).split('/')[1]);
  });
}
async function viewTopic(s, t) {
  const data = await topicData(s.id, t.id);
  const draw = () => {
    const items = t.items.filter(it => S.diff === 'All' || it.difficulty === S.diff || (S.diff === 'Todo' && !isDone(s.id, it)));
    $('#rows').innerHTML = rows(s, t.id, items);
    bindChecks($('#rows'));
  };
  main.innerHTML = `<div class="crumbs"><a href="${href(s.id)}">${esc(s.title)}</a></div><h1>${esc(t.title)}</h1><div id="intro"></div>
    ${s.id === 'dsa' ? `<div class="filters">${['All', 'Easy', 'Medium', 'Hard', 'Todo'].map(d => `<button class="${S.diff === d ? 'on' : ''}">${d}</button>`).join('')}</div>` : '<h2>Articles</h2>'}
    <div id="rows"></div>`;
  if (data.intro) $('#intro').append(md(data.intro));
  main.querySelectorAll('.filters button').forEach(b => b.onclick = () => {
    S.diff = b.textContent;
    main.querySelectorAll('.filters button').forEach(x => x.classList.toggle('on', x === b));
    draw();
  });
  draw();
}
async function viewItem(s, t, slug) {
  const data = await topicData(s.id, t.id);
  const it = data.items.find(x => x.slug === slug);
  if (!it) return notFound();
  const pos = FLAT.findIndex(f => f.s === s.id && f.it.slug === slug);
  const prev = FLAT[pos - 1]?.s === s.id ? FLAT[pos - 1] : null, next = FLAT[pos + 1]?.s === s.id ? FLAT[pos + 1] : null;
  const k = key(s.id, it), dsa = s.id === 'dsa';
  const related = dsa ? it.related.map(id => BY_ID[id]).filter(Boolean) : [];
  main.innerHTML = `
    <div class="crumbs"><a href="${href(s.id)}">${esc(s.title)}</a> › <a href="${href(s.id, t.id)}">${esc(t.title)}</a></div>
    <h1>${dsa ? it.id + '. ' : ''}${esc(it.title)}</h1>
    <div class="meta">
      ${dsa ? `<span class="badge ${it.difficulty}">${it.difficulty}</span><a href="https://leetcode.com/problems/${it.slug}/" target="_blank" rel="noopener">LeetCode ↗</a>` : ''}
      ${tags(it.tags)}
      <label class="done"><input type="checkbox" data-k="${k}" ${S.done[k] ? 'checked' : ''}> Done</label>
    </div>
    <div id="body"></div>
    ${dsa ? `<div class="code">
      <div class="bar"><button data-l="python">Python</button><button data-l="java">Java</button><span class="sp"></span>
        <button id="tests">Tests</button><button id="copy">Copy</button><button id="file">Run locally ↗</button></div>
      <pre><code id="src"></code></pre><pre class="test" hidden><code id="tsrc"></code></pre></div>` : ''}
    ${related.length ? `<h3>Related</h3><div class="chips">${related.map(f =>
      `<a href="${href(f.s, f.t, f.it.slug)}">${f.it.id}. ${esc(f.it.title)} <span class="${f.it.difficulty}">●</span></a>`).join('')}</div>` : ''}
    <div class="pager"><span>${prev ? `<a href="${href(prev.s, prev.t, prev.it.slug)}">← ${esc(prev.it.title)}</a>` : ''}</span>
      <span>${next ? `<a href="${href(next.s, next.t, next.it.slug)}">${esc(next.it.title)} →</a>` : ''}</span></div>`;
  $('#body').append(md(it.md));
  bindChecks(main);
  if (!dsa) return;
  const base = String(it.id).padStart(4, '0') + '_' + it.slug.replace(/-/g, '_');
  const cls = 'P' + String(it.id).padStart(4, '0') + it.slug.split('-').map(w => w[0].toUpperCase() + w.slice(1)).join('');
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
function viewSearch(q) {
  const words = q.toLowerCase().split(/\s+/).filter(Boolean);
  const hits = FLAT.filter(f => {
    const hay = `${f.it.id || ''} ${f.it.title} ${(f.it.tags || []).join(' ')} ${f.tTitle} ${f.it.difficulty || ''}`.toLowerCase();
    return words.every(w => hay.includes(w));
  }).slice(0, 100);
  main.innerHTML = `<h1>Search</h1><p class="muted">${hits.length} result${hits.length === 1 ? '' : 's'} for “${esc(q)}”</p>
    <table class="list">${hits.map(f => `<tr><td class="n">${f.it.id || ''}</td><td><a href="${href(f.s, f.t, f.it.slug)}">${esc(f.it.title)}</a>
    <span class="tags"><span class="tag">${esc(f.tTitle)}</span></span></td><td class="d ${f.it.difficulty || ''}">${f.it.difficulty || sec(f.s).title}</td></tr>`).join('')}</table>`;
}
function notFound() { main.innerHTML = `<h1>Not found</h1><p><a href="#/">Go home</a></p>`; }

// ---------------------------------------------------------------- router
async function route() {
  document.body.classList.remove('open');
  const [sid, tid, slug] = location.hash.replace(/^#\/?/, '').split('/').map(decodeURIComponent);
  const s = sid && sec(sid), t = s && tid && s.topics.find(x => x.id === tid);
  renderNav(s ? s.id : '');
  renderSide(s ? s.id : 'dsa', t && t.id);
  try {
    if (!sid) viewHome();
    else if (!s) notFound();
    else if (!tid) viewSection(s);
    else if (!t) notFound();
    else if (!slug) await viewTopic(s, t);
    else await viewItem(s, t, slug);
  } catch (e) { main.innerHTML = `<h1>Error</h1><pre>${esc(e.message)}</pre>`; }
  document.title = (main.querySelector('h1')?.textContent || '') + ' · LeetCode Review';
  window.scrollTo(0, 0);
}

let qTimer;
$('#q').addEventListener('input', e => {
  clearTimeout(qTimer);
  qTimer = setTimeout(() => e.target.value.trim() ? viewSearch(e.target.value.trim()) : route(), 120);
});
document.addEventListener('keydown', e => {
  if (e.key === '/' && document.activeElement !== $('#q')) { e.preventDefault(); $('#q').focus(); }
  if (e.key === 'Escape' && document.activeElement === $('#q')) { $('#q').value = ''; $('#q').blur(); route(); }
});
$('#menu').onclick = () => document.body.classList.toggle('open');
$('#theme').onclick = () => {
  const dark = matchMedia('(prefers-color-scheme: dark)').matches;
  const cur = document.documentElement.dataset.theme || (dark ? 'dark' : 'light');
  const nxt = cur === 'dark' ? 'light' : 'dark';
  document.documentElement.dataset.theme = nxt;
  store.set('lc-theme', nxt);
};
window.addEventListener('hashchange', () => { $('#q').value = ''; route(); });

getJSON('data/manifest.json').then(m => { S.m = m; index(); route(); })
  .catch(e => { main.innerHTML = `<h1>Could not load data</h1><pre>${esc(e.message)}</pre>`; });
