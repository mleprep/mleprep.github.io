#!/usr/bin/env python3
"""Build data/companies.json from a local clone of the public dataset
https://github.com/liquidslr/leetcode-company-wise-problems (one folder per company,
CSV files per time window with columns Difficulty,Title,Frequency,Acceptance Rate,Link,Topics).

    git clone --depth 1 https://github.com/liquidslr/leetcode-company-wise-problems.git /some/dir
    python tools/import_companies.py /some/dir

Output (only problems that exist in content/dsa):
    {"source": {...}, "companies": [names...],
     "problems": {"two-sum": [[companyIndex, freqAllTime, freqLast6Months], ...]}}   # sorted by freqAllTime desc
Frequency is the dataset's 0-100 score, relative within each company (100 = that company's most-asked problem).
"""
import csv
import json
import re
import subprocess
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
WINDOWS = {"all": "5. All.csv", "6m": "3. Six Months.csv"}


def our_slugs():
    m = json.loads((ROOT / "data" / "manifest.json").read_text(encoding="utf-8"))
    dsa = next(s for s in m["sections"] if s["id"] == "dsa")
    return {it["slug"] for t in dsa["topics"] for it in t["items"]}


def read(path):
    out = {}
    if not path.exists():
        return out
    with path.open(encoding="utf-8", newline="") as f:
        for row in csv.DictReader(f):
            m = re.search(r"/problems/([^/?#]+)", row.get("Link") or "")
            try:
                freq = float(row.get("Frequency") or 0)
            except ValueError:
                continue
            if m:
                out[m.group(1).lower()] = freq
    return out


def main():
    if len(sys.argv) < 2:
        sys.exit(__doc__)
    src = Path(sys.argv[1])
    slugs = our_slugs()
    per = {}   # slug -> {company: [all, 6m]}
    names = sorted((d.name for d in src.iterdir() if d.is_dir() and not d.name.startswith(".")), key=str.lower)
    for name in names:
        allf, recent = read(src / name / WINDOWS["all"]), read(src / name / WINDOWS["6m"])
        for slug in (set(allf) | set(recent)) & slugs:
            per.setdefault(slug, {})[name] = [round(allf.get(slug, 0), 1), round(recent.get(slug, 0), 1)]
    used = sorted({c for v in per.values() for c in v}, key=str.lower)
    idx = {c: i for i, c in enumerate(used)}
    problems = {s: sorted(([idx[c], a, r] for c, (a, r) in v.items()), key=lambda x: (-x[1], -x[2], used[x[0]]))
                for s, v in sorted(per.items())}
    try:
        commit = subprocess.run(["git", "-C", str(src), "log", "-1", "--format=%h %cs"],
                                capture_output=True, text=True).stdout.split()
    except OSError:
        commit = []
    out = {"source": {"repo": "https://github.com/liquidslr/leetcode-company-wise-problems",
                      "commit": commit[0] if commit else "", "date": commit[1] if len(commit) > 1 else ""},
           "companies": used, "problems": problems}
    (ROOT / "data" / "companies.json").write_text(json.dumps(out, ensure_ascii=False, separators=(",", ":")),
                                                  encoding="utf-8")
    missing = sorted(slugs - set(problems))
    print(f"{len(problems)}/{len(slugs)} problems tagged, {len(used)} companies, "
          f"{sum(len(v) for v in problems.values())} problem-company pairs")
    if missing:
        print("no company data:", ", ".join(missing))


if __name__ == "__main__":
    main()
