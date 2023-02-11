from http.server import BaseHTTPRequestHandler, HTTPServer

import glob
import json

path = r'/app/json/*.json'


def getUrlAndResponseList(path):
    urlAndResponse = []
    print("getUrlAndResponseList start")
    files = glob.glob(path)
    for i in files:
        f = open(i)

        data = json.load(f)
        request = {}
        request["url"] = data["url"]
        request["response"] = data["response"]
        print(data)
        
        urlAndResponse.append(request)

    return urlAndResponse


PORT = 8080
hostName = "0.0.0.0"
serverPort = 8080

class MyServer(BaseHTTPRequestHandler):
    def do_GET(self):

        
        self.send_response(200)
        
        self.send_header('Access-Control-Allow-Origin', '*')
        self.send_header("Content-type", "application/json")
        self.end_headers()
        objectWithUrlOfRequest = list(filter(lambda urlAndResponse: urlAndResponse['url'] == self.path, getUrlAndResponseList(path)))
        
        if len(objectWithUrlOfRequest) > 0:
            self.wfile.write(bytes(json.dumps(objectWithUrlOfRequest[0].get('response')), "utf-8"))

        

   
webServer = HTTPServer((hostName, serverPort), MyServer)
print("Server started http://%s:%s" % (hostName, serverPort))

try:
    webServer.serve_forever()
except KeyboardInterrupt:
    pass

webServer.server_close()
print("Server stopped.")

    