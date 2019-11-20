import pytest


def test_uno(default_worker):
    clientId = "123456789"
    request = {"name":"dummy"}
    response = default_worker.create(clientId, request)

    assert response == 200

