# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import test_pb2 as test__pb2


class TestStub(object):
  """定义服务
  """

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.SayHello = channel.unary_unary(
        '/Test/SayHello',
        request_serializer=test__pb2.MyRequest.SerializeToString,
        response_deserializer=test__pb2.MyReply.FromString,
        )


class TestServicer(object):
  """定义服务
  """

  def SayHello(self, request, context):
    """在服务中定义接口(指定请求和相应类型)
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_TestServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'SayHello': grpc.unary_unary_rpc_method_handler(
          servicer.SayHello,
          request_deserializer=test__pb2.MyRequest.FromString,
          response_serializer=test__pb2.MyReply.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'Test', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))