// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.12.4
// source: proto/quiz.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NoRequestParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoRequestParam) Reset() {
	*x = NoRequestParam{}
	mi := &file_proto_quiz_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NoRequestParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoRequestParam) ProtoMessage() {}

func (x *NoRequestParam) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quiz_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoRequestParam.ProtoReflect.Descriptor instead.
func (*NoRequestParam) Descriptor() ([]byte, []int) {
	return file_proto_quiz_proto_rawDescGZIP(), []int{0}
}

type Question struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	QuestionDesc string   `protobuf:"bytes,2,opt,name=question_desc,json=questionDesc,proto3" json:"question_desc,omitempty"`
	Options      []string `protobuf:"bytes,3,rep,name=options,proto3" json:"options,omitempty"`
}

func (x *Question) Reset() {
	*x = Question{}
	mi := &file_proto_quiz_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Question) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Question) ProtoMessage() {}

func (x *Question) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quiz_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Question.ProtoReflect.Descriptor instead.
func (*Question) Descriptor() ([]byte, []int) {
	return file_proto_quiz_proto_rawDescGZIP(), []int{1}
}

func (x *Question) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Question) GetQuestionDesc() string {
	if x != nil {
		return x.QuestionDesc
	}
	return ""
}

func (x *Question) GetOptions() []string {
	if x != nil {
		return x.Options
	}
	return nil
}

type QuestionsResponsePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Questions []*Question `protobuf:"bytes,1,rep,name=questions,proto3" json:"questions,omitempty"`
}

func (x *QuestionsResponsePayload) Reset() {
	*x = QuestionsResponsePayload{}
	mi := &file_proto_quiz_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QuestionsResponsePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionsResponsePayload) ProtoMessage() {}

func (x *QuestionsResponsePayload) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quiz_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionsResponsePayload.ProtoReflect.Descriptor instead.
func (*QuestionsResponsePayload) Descriptor() ([]byte, []int) {
	return file_proto_quiz_proto_rawDescGZIP(), []int{2}
}

func (x *QuestionsResponsePayload) GetQuestions() []*Question {
	if x != nil {
		return x.Questions
	}
	return nil
}

type Answer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestionId int32  `protobuf:"varint,1,opt,name=question_id,json=questionId,proto3" json:"question_id,omitempty"`
	Answer     string `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *Answer) Reset() {
	*x = Answer{}
	mi := &file_proto_quiz_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Answer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Answer) ProtoMessage() {}

func (x *Answer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quiz_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Answer.ProtoReflect.Descriptor instead.
func (*Answer) Descriptor() ([]byte, []int) {
	return file_proto_quiz_proto_rawDescGZIP(), []int{3}
}

func (x *Answer) GetQuestionId() int32 {
	if x != nil {
		return x.QuestionId
	}
	return 0
}

func (x *Answer) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

type AnswersRequestPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string    `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Answers []*Answer `protobuf:"bytes,2,rep,name=answers,proto3" json:"answers,omitempty"`
}

func (x *AnswersRequestPayload) Reset() {
	*x = AnswersRequestPayload{}
	mi := &file_proto_quiz_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnswersRequestPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnswersRequestPayload) ProtoMessage() {}

func (x *AnswersRequestPayload) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quiz_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnswersRequestPayload.ProtoReflect.Descriptor instead.
func (*AnswersRequestPayload) Descriptor() ([]byte, []int) {
	return file_proto_quiz_proto_rawDescGZIP(), []int{4}
}

func (x *AnswersRequestPayload) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AnswersRequestPayload) GetAnswers() []*Answer {
	if x != nil {
		return x.Answers
	}
	return nil
}

type ResultResponsePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrectAnswerCount int32   `protobuf:"varint,1,opt,name=correct_answer_count,json=correctAnswerCount,proto3" json:"correct_answer_count,omitempty"`
	ResultPercentage   float32 `protobuf:"fixed32,2,opt,name=result_percentage,json=resultPercentage,proto3" json:"result_percentage,omitempty"`
	UserId             string  `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ResultResponsePayload) Reset() {
	*x = ResultResponsePayload{}
	mi := &file_proto_quiz_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResultResponsePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultResponsePayload) ProtoMessage() {}

func (x *ResultResponsePayload) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quiz_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultResponsePayload.ProtoReflect.Descriptor instead.
func (*ResultResponsePayload) Descriptor() ([]byte, []int) {
	return file_proto_quiz_proto_rawDescGZIP(), []int{5}
}

func (x *ResultResponsePayload) GetCorrectAnswerCount() int32 {
	if x != nil {
		return x.CorrectAnswerCount
	}
	return 0
}

func (x *ResultResponsePayload) GetResultPercentage() float32 {
	if x != nil {
		return x.ResultPercentage
	}
	return 0
}

func (x *ResultResponsePayload) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ResultList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*ResultResponsePayload `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ResultList) Reset() {
	*x = ResultList{}
	mi := &file_proto_quiz_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResultList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultList) ProtoMessage() {}

func (x *ResultList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quiz_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultList.ProtoReflect.Descriptor instead.
func (*ResultList) Descriptor() ([]byte, []int) {
	return file_proto_quiz_proto_rawDescGZIP(), []int{6}
}

func (x *ResultList) GetResults() []*ResultResponsePayload {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_proto_quiz_proto protoreflect.FileDescriptor

var file_proto_quiz_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x71, 0x75, 0x69, 0x7a, 0x22, 0x10, 0x0a, 0x0e, 0x4e, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x22, 0x59, 0x0a, 0x08, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x48, 0x0a, 0x18, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x2c, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x41, 0x0a, 0x06, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x22, 0x58, 0x0a, 0x15, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x41, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x52, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x22, 0x8f, 0x01, 0x0a,
	0x15, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63,
	0x74, 0x5f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x41, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x10, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x43,
	0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x71, 0x75, 0x69, 0x7a, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x32, 0xa1, 0x01, 0x0a, 0x0e, 0x43, 0x4c, 0x49, 0x51, 0x75, 0x69, 0x7a, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x4e, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x1e, 0x2e, 0x71,
	0x75, 0x69, 0x7a, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x49, 0x0a, 0x0d,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x12, 0x1b, 0x2e,
	0x71, 0x75, 0x69, 0x7a, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x1b, 0x2e, 0x71, 0x75, 0x69,
	0x7a, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_quiz_proto_rawDescOnce sync.Once
	file_proto_quiz_proto_rawDescData = file_proto_quiz_proto_rawDesc
)

func file_proto_quiz_proto_rawDescGZIP() []byte {
	file_proto_quiz_proto_rawDescOnce.Do(func() {
		file_proto_quiz_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_quiz_proto_rawDescData)
	})
	return file_proto_quiz_proto_rawDescData
}

var file_proto_quiz_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_quiz_proto_goTypes = []any{
	(*NoRequestParam)(nil),           // 0: quiz.NoRequestParam
	(*Question)(nil),                 // 1: quiz.Question
	(*QuestionsResponsePayload)(nil), // 2: quiz.QuestionsResponsePayload
	(*Answer)(nil),                   // 3: quiz.Answer
	(*AnswersRequestPayload)(nil),    // 4: quiz.AnswersRequestPayload
	(*ResultResponsePayload)(nil),    // 5: quiz.ResultResponsePayload
	(*ResultList)(nil),               // 6: quiz.ResultList
}
var file_proto_quiz_proto_depIdxs = []int32{
	1, // 0: quiz.QuestionsResponsePayload.questions:type_name -> quiz.Question
	3, // 1: quiz.AnswersRequestPayload.answers:type_name -> quiz.Answer
	5, // 2: quiz.ResultList.results:type_name -> quiz.ResultResponsePayload
	0, // 3: quiz.CLIQuizService.GetQuestions:input_type -> quiz.NoRequestParam
	4, // 4: quiz.CLIQuizService.SubmitAnswers:input_type -> quiz.AnswersRequestPayload
	2, // 5: quiz.CLIQuizService.GetQuestions:output_type -> quiz.QuestionsResponsePayload
	5, // 6: quiz.CLIQuizService.SubmitAnswers:output_type -> quiz.ResultResponsePayload
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_quiz_proto_init() }
func file_proto_quiz_proto_init() {
	if File_proto_quiz_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_quiz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_quiz_proto_goTypes,
		DependencyIndexes: file_proto_quiz_proto_depIdxs,
		MessageInfos:      file_proto_quiz_proto_msgTypes,
	}.Build()
	File_proto_quiz_proto = out.File
	file_proto_quiz_proto_rawDesc = nil
	file_proto_quiz_proto_goTypes = nil
	file_proto_quiz_proto_depIdxs = nil
}
