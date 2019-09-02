from concurrent import futures
import os
import time
import math
import logging

import json

import grpc
import protolib
import socket

from tensorflow.keras.models import model_from_json
import numpy as np

ONE_DAY_IN_SECONDS = 60 * 60 * 24


class RouteGuideServicer:
    def __init__(self):
        # In tensorflow 1.14.0, Got error: call initializer instance with the dtype argument instead of passing it to the constructor.
        # Need upgrade to tensorflow 2.0.0rc0 to avoid it.
        super().__init__()
        with open("./models/simple_nn.json", "r") as f:
            json_string = json.load(f)

        self.model = model_from_json(json_string)
        self.model.load_weights("./models/simple_nn_weights.h5")

    def Predict(self, request, context):
        # TODO Irisのコードを差し込む
        input = [
            request.sepalLength,
            request.sepalWidth,
            request.petalLength,
            request.petalWidth,
        ]

        output = np.argmax(self.model.predict(np.array([input])))
        labels = ["Iris-setosa", "Iris-versicolor", "Iris-virginica"]

        # host = socket.gethostname()
        # ip = socket.gethostbyname(host)
        return protolib.Response(irisType=labels[output])


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    protolib.add_AdTechServicer_to_server(RouteGuideServicer(), server)

    port = "50051"
    if os.environ.get("GRPC_PORT") != None:
        port = os.environ.get("GRPC_PORT")

    server.add_insecure_port("[::]:" + port)
    server.start()
    try:
        while True:
            time.sleep(ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == "__main__":
    logging.basicConfig()
    serve()
