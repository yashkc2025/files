url='https://www.google.pl'

if avg_speed=$(curl -sS -H 'Cache-Control: no-store' -w '%{http_code} %{time_total}' -o /dev/null --url "$url")
then
  echo "$avg_speed"
fi
