
SERVER_URL=127.0.0.1:8080

FOLDER=tmp
mkdir -p $FOLDER

# MOVE_LEFT=0
# MOVE_UP=1
# MOVE_RIGHT=2
# MOVE_DOWN=3

PRIMARY_KEY=1
# MOVE=1
GODMODE=true

BODY=$(cat << EOM
{"primarykey":$PRIMARY_KEY, 
"move":$1,
"godmode":$2
}
EOM
)

#put submit
echo "MOVE/"
curl -X PUT "${SERVER_URL}/move" -H "Content-Type: application/json" --data "${BODY}//$'\n'/" | jq '.' | jq -MRsr 'gsub("\n      +";"")|gsub("\n    ]";"]")'> $FOLDER/move.json
echo