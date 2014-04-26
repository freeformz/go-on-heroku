hk create go-on-heroku // HL
hk set BUILDPACK_URL=https://github.com/kr/heroku-buildpack-go.git

git push heroku master // HL

hk env
hk dynos
hk log -s app
hk open
