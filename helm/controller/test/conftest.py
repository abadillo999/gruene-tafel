import pytest

import http.client
import json

class http_client:

    def __init__(self, url, port):
        self.http_connection = http.client.HTTPConnection(url, port)
        yield
        self.http_connection.close()

    def post(self, uri, headers, data):
        self.http_connection.request('POST', uri, data, headers)
        return self.http_connection.getresponse()

    def get(self, uri):
        self.http_connection.request('GET', uri)
        return self.http_connection.getresponse()

    def delete(self, uri):
        self.http_connection.request('DELETE', uri)
        return self.http_connection.getresponse()


class worker_client:
    client = None
    headers = {"Content-type": "application/json"}

    def start(self, url, port):
        self.client = http_client(url, port)

    def create(self, clientid, data):
        #TODO: Add clientId to data
        json_data = json.dumps(data, clientid=clientId)
        return self.client.post("/task", self.headers, json_data )

    def update(self, clientId, data):
        json_data = json.dumps(data)
        return self.client.post("/task/{}".format(clientId), json_data)

    def delete(self, clientId, data):
        return self.client.delete("/task/{}".format(clientId))



@pytest.fixture(scope="module")
def default_worker():
    client = worker_client()
    client.start("default-worker-service", 8080)
    yield client
