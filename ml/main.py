from concurrent import futures
import os
import time
import math
import logging
import json
import protolib
import socket
import numpy as np
import pickle

# from tensorflow.keras.models import model_from_json
import grpc

ONE_DAY_IN_SECONDS = 60 * 60 * 24


def load_svm():
    with open("./models/svm.pkl", "rb") as file:
        model = pickle.load(file)
    return model

class RouteGuideServicer:
    def __init__(self):
        # In tensorflow 1.14.0, Got error: call initializer instance with the dtype argument instead of passing it to the constructor.
        # Need upgrade to tensorflow 2.0.0rc0 to avoid it.
        super().__init__()
        self.model = load_svm()

    def Predict(self, request, context):
        input = [
            request.sepalLength,
            request.sepalWidth,
            request.petalLength,
            request.petalWidth,
        ]
        output = self.model.predict(np.array([input]))[0]
        labels = ["Iris-setosa", "Iris-versicolor", "Iris-virginica"]

        host = socket.gethostname()
        ip = socket.gethostbyname(host)
        return protolib.Response(irisType=labels[output] + "(" + host +  ":" + ip + ")")


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    protolib.add_ProtoServicer_to_server(RouteGuideServicer(), server)

    port = "50051"
    if os.environ.get("GRPC_PORT") != None:
        port = os.environ.get("GRPC_PORT")

    server.add_insecure_port("[::]:" + port)
    server.start()
    print("server running on port:", port)
    try:
        while True:
            time.sleep(ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == "__main__":
    logging.basicConfig()
    serve()
