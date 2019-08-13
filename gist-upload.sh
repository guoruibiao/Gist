# 0. Your file name
FNAME=filename

CONTENT=helloworldthishidsidhsidhsids

# 2. Build the JSON request
read -r -d '' DESC <<EOF
{
  "description": "some description",
  "public": true,
  "files": {
    "helloworld": {
      "content": "${CONTENT}"
    }
  }
}
EOF

# 3. Use curl to send a POST request
#curl -u "username" -X POST -d "${DESC}" "https://api.github.com/gists"
curl -H "Authorization: token xxx" -X POST -d "${DESC}" "https://api.github.com/gists"