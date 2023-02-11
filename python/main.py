from http.server import BaseHTTPRequestHandler, HTTPServer

import glob
import json

urlAndResponse = []

path = r'/app/json/*.json'


def getUrlAndResponseList(path):
    print("getUrlAndResponseList start")
    files = glob.glob(path)
    for i in files:
        f = open(i)

        data = json.load(f)
        request = {}
        request["url"] = data["url"]
        request["response"] = data["response"]
        
        urlAndResponse.append(request)

    return urlAndResponse

urlAndResponse = getUrlAndResponseList(path)

PORT = 8080
hostName = "0.0.0.0"
serverPort = 8080

class MyServer(BaseHTTPRequestHandler):
    def do_GET(self):

        
        self.send_response(200)
        
        self.send_header('Access-Control-Allow-Origin', '*')
        self.send_header("Content-type", "application/json")
        self.end_headers()
        urlAndResponse1 = getUrlAndResponseList(path)
        ss = filter(lambda person: person['url'] == self.path, urlAndResponse1)
        ss2 = list(ss)
        if len(ss2) > 0:
            self.wfile.write(bytes(json.dumps(ss2[0].get('response')), "utf-8"))

        

   
webServer = HTTPServer((hostName, serverPort), MyServer)
print("Server started http://%s:%s" % (hostName, serverPort))

try:
    webServer.serve_forever()
except KeyboardInterrupt:
    pass

webServer.server_close()
print("Server stopped.")

    