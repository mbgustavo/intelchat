# Build backend
cd ../backend
# Executable main file
mkdir -p build
go build -o build/
# Copy production configuration json file to the same build folder
mkdir -p build/configs
cp configs/configs.prod.json build/configs/configs.json

# Build frontend
cd ../frontend
npm run build