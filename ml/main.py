from concurrent import futures
import time
import math
import logging

import grpc
import protolib

ONE_DAY_IN_SECONDS = 60 * 60 * 24

class RouteGuideServicer():

    def Predict(self, request, context):
        print(request)
        print(context)
        return protolib.Response(irisType="how do i know with such little info")


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    protolib.add_AdTechServicer_to_server(
        RouteGuideServicer(),
        server
    )
    server.add_insecure_port('[::]:50051')
    server.start()
    try:
        while True:
            time.sleep(ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    logging.basicConfig()
    serve()
