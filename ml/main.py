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
import pickle

ONE_DAY_IN_SECONDS = 60 * 60 * 24


def load_svm():
    with open("./models/svm.pkl", "rb") as file:
        model = pickle.load(file)
    return model


def load_nn():
    with open("./models/simple_nn.json", "r") as f:
        json_string = json.load(f)

    model = model_from_json(json_string)
    model.load_weights("./models/simple_nn_weights.h5")
    return model


class RouteGuideServicer:
    def __init__(self):
        # In tensorflow 1.14.0, Got error: call initializer instance with the dtype argument instead of passing it to the constructor.
        # Need upgrade to tensorflow 2.0.0rc0 to avoid it.
        super().__init__()

        self.model_type = "nn"
        if os.environ.get("MODEL_TYPE") != None:
            self.model_type = os.environ.get("MODEL_TYPE")

        if self.model_type == "nn":
            self.model = load_nn()
        elif self.model_type == "svm":
            self.model = load_svm()
        else:
            self.model = load_svm()

    def certainPredict(self, input):
        if self.model_type == "nn":
            return np.argmax(self.model.predict(np.array([input])))
        elif self.model_type == "svm":
            return self.model.predict(np.array([input]))[0]
        else:
            return np.argmax(self.model.predict(np.array([input])))

    def Predict(self, request, context):
        input = [
            request.sepalLength,
            request.sepalWidth,
            request.petalLength,
            request.petalWidth,
        ]
        output = self.certainPredict(input)
        labels = ["Iris-setosa", "Iris-versicolor", "Iris-virginica"]

        model_type = "nn"
        if os.environ.get("MODEL_TYPE") != None:
            model_type = os.environ.get("MODEL_TYPE")

        # host = socket.gethostname()
        # ip = socket.gethostbyname(host)
        return protolib.Response(irisType=labels[output] + "(" + self.model_type + ")")


def serve():
    # TODO max_workersの適切な値は?
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
    # TODO ここいじったらいい感じにログ出せる？
    logging.basicConfig()
    serve()
