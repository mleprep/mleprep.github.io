/* In-browser Python judge: Pyodide (CPython compiled to WebAssembly) in a Web Worker.
   Nothing is sent to a server — code runs on the visitor's machine. The page terminates
   this worker on a time limit (infinite loops) and starts a fresh one for the next run.
   Messages in:  {id, kind: 'judge', code, test} | {id, kind: 'script', code}
   Messages out: {ready, error?} once; {id, started} when a run begins; {id, ok, result|error} */
const PYODIDE = 'https://cdn.jsdelivr.net/pyodide/v0.29.5/full/';
importScripts(PYODIDE + 'pyodide.js');

const text = p => fetch(new URL(p, self.location)).then(r => {
  if (!r.ok) throw new Error(`${p}: HTTP ${r.status}`);
  return r.text();
});

const boot = (async () => {
  const [py, lc, judge] = await Promise.all([loadPyodide({ indexURL: PYODIDE }), text('../python/lc.py'), text('judge.py')]);
  py.FS.mkdirTree('/judge');
  py.FS.writeFile('/judge/lc.py', lc);
  py.FS.writeFile('/judge/judge.py', judge);
  py.runPython("import sys; sys.path.insert(0, '/judge'); import judge");
  return py;
})();
boot.then(() => self.postMessage({ ready: true }), e => self.postMessage({ ready: false, error: String(e && e.message || e) }));

self.onmessage = async ({ data: { id, kind, code, test } }) => {
  try {
    const py = await boot;
    await py.loadPackagesFromImports(code);   // e.g. numpy for the ML snippets (downloaded once, then cached)
    self.postMessage({ id, started: true });
    py.globals.set('_code', code);
    py.globals.set('_test', test || '');
    const res = py.runPython(kind === 'script' ? 'judge.run_script(_code)' : 'judge.judge(_code, _test)');
    self.postMessage({ id, ok: true, result: JSON.parse(res) });
  } catch (e) {
    self.postMessage({ id, ok: false, error: String(e && e.message || e) });
  }
};
