
SERVER_URL=127.0.0.1:5442

FOLDER=tmp
mkdir -p $FOLDER

USER_ID=66666666-6666-6666-6666-666666666666
LEVEL=[[1,1,1,1,0,1,1,1],[1,0,0,0,0,0,0,1],[1,0,1,1,1,3,1,1],[1,0,0,0,1,0,2,1],[1,1,1,0,1,1,0,1],[1,0,0,0,1,0,0,1],[1,0,1,1,1,0,1,1],[1,0,0,4,0,0,0,1],[1,1,1,1,1,1,1,1]]

BODY=$(cat << EOM
{"userid":"$USER_ID", 
"level":$LEVEL
}
EOM
)

#put submit
echo "SUBMIT/"
curl -X PUT "${SERVER_URL}/submit" -H "Content-Type: application/json" --data "${BODY}//$'\n'/" | jq '.' | jq -MRsr 'gsub("\n      +";"")|gsub("\n    ]";"]")'> $FOLDER/submit.json
echo