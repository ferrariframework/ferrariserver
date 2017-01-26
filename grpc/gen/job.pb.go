// Code generated by protoc-gen-go.
// source: job.proto
// DO NOT EDIT!

/*
Package gen is a generated protocol buffer package.

It is generated from these files:
	job.proto

It has these top-level messages:
	Job
	JobResult
	Worker
	JobLog
	Empty
*/
package gen

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type JobResult_Status int32

const (
	JobResult_Sucess JobResult_Status = 0
	JobResult_Failed JobResult_Status = 1
)

var JobResult_Status_name = map[int32]string{
	0: "Sucess",
	1: "Failed",
}
var JobResult_Status_value = map[string]int32{
	"Sucess": 0,
	"Failed": 1,
}

func (x JobResult_Status) String() string {
	return proto.EnumName(JobResult_Status_name, int32(x))
}
func (JobResult_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Job struct {
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	WorkerId  string `protobuf:"bytes,2,opt,name=worker_id,json=workerId" json:"worker_id,omitempty"`
	RunId     string `protobuf:"bytes,3,opt,name=run_id,json=runId" json:"run_id,omitempty"`
	StartTime int64  `protobuf:"varint,4,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	EndTime   int64  `protobuf:"varint,5,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Job) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Job) GetWorkerId() string {
	if m != nil {
		return m.WorkerId
	}
	return ""
}

func (m *Job) GetRunId() string {
	if m != nil {
		return m.RunId
	}
	return ""
}

func (m *Job) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Job) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

type JobResult struct {
	WorkerId  string `protobuf:"bytes,1,opt,name=worker_id,json=workerId" json:"worker_id,omitempty"`
	JobId     string `protobuf:"bytes,2,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	RungId    string `protobuf:"bytes,3,opt,name=rung_id,json=rungId" json:"rung_id,omitempty"`
	StartTime int64  `protobuf:"varint,4,opt,name=startTime" json:"startTime,omitempty"`
}

func (m *JobResult) Reset()                    { *m = JobResult{} }
func (m *JobResult) String() string            { return proto.CompactTextString(m) }
func (*JobResult) ProtoMessage()               {}
func (*JobResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *JobResult) GetWorkerId() string {
	if m != nil {
		return m.WorkerId
	}
	return ""
}

func (m *JobResult) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *JobResult) GetRungId() string {
	if m != nil {
		return m.RungId
	}
	return ""
}

func (m *JobResult) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

type Worker struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *Worker) Reset()                    { *m = Worker{} }
func (m *Worker) String() string            { return proto.CompactTextString(m) }
func (*Worker) ProtoMessage()               {}
func (*Worker) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Worker) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type JobLog struct {
	WorkerId string `protobuf:"bytes,1,opt,name=worker_id,json=workerId" json:"worker_id,omitempty"`
	RunId    string `protobuf:"bytes,2,opt,name=run_id,json=runId" json:"run_id,omitempty"`
	JobId    string `protobuf:"bytes,3,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	Message  []byte `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *JobLog) Reset()                    { *m = JobLog{} }
func (m *JobLog) String() string            { return proto.CompactTextString(m) }
func (*JobLog) ProtoMessage()               {}
func (*JobLog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *JobLog) GetWorkerId() string {
	if m != nil {
		return m.WorkerId
	}
	return ""
}

func (m *JobLog) GetRunId() string {
	if m != nil {
		return m.RunId
	}
	return ""
}

func (m *JobLog) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *JobLog) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*Job)(nil), "gen.Job")
	proto.RegisterType((*JobResult)(nil), "gen.JobResult")
	proto.RegisterType((*Worker)(nil), "gen.Worker")
	proto.RegisterType((*JobLog)(nil), "gen.JobLog")
	proto.RegisterType((*Empty)(nil), "gen.Empty")
	proto.RegisterEnum("gen.JobResult_Status", JobResult_Status_name, JobResult_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for JobService service

type JobServiceClient interface {
	// RegisterJob registers a job given a Worker.ID message
	RegisterJob(ctx context.Context, in *Job, opts ...grpc.CallOption) (*Job, error)
	// RecordLog records
	RecordLog(ctx context.Context, opts ...grpc.CallOption) (JobService_RecordLogClient, error)
	RegisterJobResult(ctx context.Context, in *JobResult, opts ...grpc.CallOption) (*Job, error)
}

type jobServiceClient struct {
	cc *grpc.ClientConn
}

func NewJobServiceClient(cc *grpc.ClientConn) JobServiceClient {
	return &jobServiceClient{cc}
}

func (c *jobServiceClient) RegisterJob(ctx context.Context, in *Job, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := grpc.Invoke(ctx, "/gen.JobService/RegisterJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) RecordLog(ctx context.Context, opts ...grpc.CallOption) (JobService_RecordLogClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_JobService_serviceDesc.Streams[0], c.cc, "/gen.JobService/RecordLog", opts...)
	if err != nil {
		return nil, err
	}
	x := &jobServiceRecordLogClient{stream}
	return x, nil
}

type JobService_RecordLogClient interface {
	Send(*JobLog) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type jobServiceRecordLogClient struct {
	grpc.ClientStream
}

func (x *jobServiceRecordLogClient) Send(m *JobLog) error {
	return x.ClientStream.SendMsg(m)
}

func (x *jobServiceRecordLogClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *jobServiceClient) RegisterJobResult(ctx context.Context, in *JobResult, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := grpc.Invoke(ctx, "/gen.JobService/RegisterJobResult", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for JobService service

type JobServiceServer interface {
	// RegisterJob registers a job given a Worker.ID message
	RegisterJob(context.Context, *Job) (*Job, error)
	// RecordLog records
	RecordLog(JobService_RecordLogServer) error
	RegisterJobResult(context.Context, *JobResult) (*Job, error)
}

func RegisterJobServiceServer(s *grpc.Server, srv JobServiceServer) {
	s.RegisterService(&_JobService_serviceDesc, srv)
}

func _JobService_RegisterJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).RegisterJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gen.JobService/RegisterJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).RegisterJob(ctx, req.(*Job))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_RecordLog_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(JobServiceServer).RecordLog(&jobServiceRecordLogServer{stream})
}

type JobService_RecordLogServer interface {
	SendAndClose(*Empty) error
	Recv() (*JobLog, error)
	grpc.ServerStream
}

type jobServiceRecordLogServer struct {
	grpc.ServerStream
}

func (x *jobServiceRecordLogServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *jobServiceRecordLogServer) Recv() (*JobLog, error) {
	m := new(JobLog)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _JobService_RegisterJobResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).RegisterJobResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gen.JobService/RegisterJobResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).RegisterJobResult(ctx, req.(*JobResult))
	}
	return interceptor(ctx, in, info, handler)
}

var _JobService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gen.JobService",
	HandlerType: (*JobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterJob",
			Handler:    _JobService_RegisterJob_Handler,
		},
		{
			MethodName: "RegisterJobResult",
			Handler:    _JobService_RegisterJobResult_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RecordLog",
			Handler:       _JobService_RecordLog_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "job.proto",
}

func init() { proto.RegisterFile("job.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x92, 0xc1, 0xce, 0xd3, 0x30,
	0x0c, 0xc7, 0x97, 0x96, 0xb6, 0xab, 0x87, 0xa6, 0x11, 0x09, 0x51, 0x06, 0x48, 0x53, 0xb8, 0xf4,
	0x34, 0x24, 0x78, 0x06, 0x90, 0x5a, 0xed, 0xd4, 0x21, 0x71, 0x9c, 0x9a, 0xc5, 0x8a, 0x32, 0xd6,
	0x66, 0x4a, 0x52, 0x10, 0x67, 0xae, 0x3c, 0x00, 0x8f, 0x8b, 0x9a, 0x6e, 0xb4, 0xdf, 0xb7, 0xc3,
	0x77, 0xb3, 0xfd, 0x57, 0xec, 0x9f, 0xff, 0x0e, 0xa4, 0x27, 0xcd, 0xb7, 0x17, 0xa3, 0x9d, 0xa6,
	0xa1, 0xc4, 0x96, 0xfd, 0x26, 0x10, 0x96, 0x9a, 0xd3, 0x25, 0x04, 0x4a, 0x64, 0x64, 0x43, 0xf2,
	0xb4, 0x0a, 0x94, 0xa0, 0x6f, 0x20, 0xfd, 0xa9, 0xcd, 0x77, 0x34, 0x07, 0x25, 0xb2, 0xc0, 0x97,
	0xe7, 0x43, 0xa1, 0x10, 0xf4, 0x25, 0xc4, 0xa6, 0x6b, 0x7b, 0x25, 0xf4, 0x4a, 0x64, 0xba, 0xb6,
	0x10, 0xf4, 0x1d, 0x80, 0x75, 0xb5, 0x71, 0x07, 0xa7, 0x1a, 0xcc, 0x9e, 0x6d, 0x48, 0x1e, 0x56,
	0xa9, 0xaf, 0x7c, 0x55, 0x0d, 0xd2, 0xd7, 0x30, 0xc7, 0x56, 0x0c, 0x62, 0xe4, 0xc5, 0x04, 0x5b,
	0xd1, 0x4b, 0xec, 0x2f, 0x81, 0xb4, 0xd4, 0xbc, 0x42, 0xdb, 0x9d, 0xdd, 0xc3, 0xd9, 0xe4, 0x7e,
	0xf6, 0x49, 0xf3, 0x91, 0x2a, 0x3a, 0x69, 0x5e, 0x08, 0xfa, 0x0a, 0x12, 0xd3, 0xb5, 0x72, 0x64,
	0xea, 0x09, 0x65, 0x21, 0xe8, 0x5b, 0x18, 0x11, 0xee, 0x98, 0xd8, 0x06, 0xe2, 0xbd, 0xab, 0x5d,
	0x67, 0x29, 0x40, 0xbc, 0xef, 0x8e, 0x68, 0xed, 0x6a, 0xd6, 0xc7, 0x5f, 0x6a, 0x75, 0x46, 0xb1,
	0x22, 0x2c, 0x83, 0xf8, 0x9b, 0x9f, 0xfd, 0xd8, 0x22, 0xd6, 0x40, 0x5c, 0x6a, 0xbe, 0xd3, 0xf2,
	0x49, 0xe0, 0xab, 0x59, 0xc1, 0xd4, 0xac, 0x71, 0x8f, 0x70, 0xba, 0x47, 0x06, 0x49, 0x83, 0xd6,
	0xd6, 0x72, 0x80, 0x7d, 0x5e, 0xdd, 0x52, 0x96, 0x40, 0xf4, 0xb9, 0xb9, 0xb8, 0x5f, 0x1f, 0xff,
	0x10, 0x80, 0x52, 0xf3, 0x3d, 0x9a, 0x1f, 0xea, 0x88, 0xf4, 0x3d, 0x2c, 0x2a, 0x94, 0xca, 0x3a,
	0x34, 0xfd, 0x21, 0xe7, 0x5b, 0x89, 0xed, 0xb6, 0xd4, 0x7c, 0xfd, 0x3f, 0x62, 0x33, 0x9a, 0x43,
	0x5a, 0xe1, 0x51, 0x1b, 0xd1, 0xe3, 0x2e, 0x6e, 0xc2, 0x4e, 0xcb, 0x35, 0xf8, 0xc4, 0x77, 0x66,
	0xb3, 0x9c, 0xd0, 0x0f, 0xf0, 0x62, 0xd2, 0xee, 0x7a, 0x91, 0xe5, 0xed, 0xc5, 0x90, 0x4f, 0x5b,
	0xf3, 0xd8, 0xff, 0xa6, 0x4f, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x63, 0xa7, 0xb4, 0x5a,
	0x02, 0x00, 0x00,
}
