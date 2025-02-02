GOOS=$1
GOARCH=$2
OUTPUT_DIR=$3
APP_NAME=$4
DO_NOT_BUILD_FRONTEND=${5:-false}

if [ -z "$GOOS" ] || [ -z "$GOARCH" ] || [ -z "$OUTPUT_DIR" ] || [ -z "$APP_NAME" ]; then
  echo "Usage: $0 <GOOS> <GOARCH> <OUTPUT_DIR> <APP_NAME> [<DO_NOT_BUILD_FRONTEND>]"
  exit 1
fi

if [[ "$DO_NOT_BUILD_FRONTEND" == "false" ]];then
  echo "Building frontend..."
  cd web || exit 1
  npm install && npm run build
  if [ $? -ne 0 ]; then
    echo "Frontend build failed."
    exit 1
  fi
  echo "Frontend build completed."
  cd ..
  touch assets/web/.gitkeep
fi

echo "Building backend for ${GOOS}/${GOARCH}..."

mkdir -p "${OUTPUT_DIR}"


if [ "$GOOS" = "windows" ]; then
  APP_NAME="${APP_NAME}.exe"
fi
go build -o "${OUTPUT_DIR}/${APP_NAME}"

if [ $? -eq 0 ]; then
  echo "Backend build successful: ${OUTPUT_DIR}/${APP_NAME}"
else
  echo "Backend build failed."
  exit 1
fi

