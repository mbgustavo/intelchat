# Execute built application, both backend and frontend should be compiled and serveFiles must be active
cd ../backend/build
./intelchat

# If serveFiles is not active, only backend will run with the command above.
# To serve the frontend files separately, you could do that with serve, installing it with the following command:
# npm install -g serve
# To serve the frontend built files, run the following command in the frontend folder to run it in the port 5000:
# serve -s build -l 5000