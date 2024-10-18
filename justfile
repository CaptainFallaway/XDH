dev:
  wails dev

gs:
  git status

gd file=".":
  git diff {{file}}
  

gac changes message:
  git add {{changes}} && git commit -m "{{message}}"


