for f in *; do
  if [ -d "$f" ]; then
    cd $f && go test -v
    cd ..
  fi
done
