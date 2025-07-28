
SERVER_URL=127.0.0.1:5442

FOLDER=tmp
mkdir -p $FOLDER

PRIMARY_KEY=1
MOVE=0

BODY=$(cat << EOM
{"primarykey":$PRIMARY_KEY, 
"move":$MOVE
}
EOM
)

#put submit
echo "MOVE/"
curl -X PUT "${SERVER_URL}/move" -H "Content-Type: application/json" --data "${BODY}//$'\n'/" | jq '.' | jq -MRsr 'gsub("\n      +";"")|gsub("\n    ]";"]")'> $FOLDER/move.json
echo