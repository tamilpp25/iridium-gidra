// Sorapointa - A server software re-implementation for a certain anime game, and avoid sorapointa.
// Copyright (C) 2022  Sorapointa Team
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.11.3
// source: GCGDuel.proto

package gen

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

type GCGDuel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerSeq                 uint32                   `protobuf:"varint,3,opt,name=server_seq,json=serverSeq,proto3" json:"server_seq,omitempty"`
	ShowInfoList              []*GCGControllerShowInfo `protobuf:"bytes,7,rep,name=show_info_list,json=showInfoList,proto3" json:"show_info_list,omitempty"`
	ForbidFinishChallengeList []uint32                 `protobuf:"varint,192,rep,packed,name=forbid_finish_challenge_list,json=forbidFinishChallengeList,proto3" json:"forbid_finish_challenge_list,omitempty"`
	CardList                  []*GCGCard               `protobuf:"bytes,1,rep,name=card_list,json=cardList,proto3" json:"card_list,omitempty"`
	Unk3300_BIANMOPDEHO       uint32                   `protobuf:"varint,9,opt,name=Unk3300_BIANMOPDEHO,json=Unk3300BIANMOPDEHO,proto3" json:"Unk3300_BIANMOPDEHO,omitempty"`
	CostRevise                *GCGCostReviseInfo       `protobuf:"bytes,8,opt,name=cost_revise,json=costRevise,proto3" json:"cost_revise,omitempty"`
	GameId                    uint32                   `protobuf:"varint,4,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	FieldList                 []*GCGPlayerField        `protobuf:"bytes,5,rep,name=field_list,json=fieldList,proto3" json:"field_list,omitempty"`
	Unk3300_CDCMBOKBLAK       []*Unk3300_ADHENCIFKNI   `protobuf:"bytes,1987,rep,name=Unk3300_CDCMBOKBLAK,json=Unk3300CDCMBOKBLAK,proto3" json:"Unk3300_CDCMBOKBLAK,omitempty"`
	BusinessType              GCGGameBusinessType      `protobuf:"varint,13,opt,name=business_type,json=businessType,proto3,enum=GCGGameBusinessType" json:"business_type,omitempty"`
	IntentionList             []*GCGPVEIntention       `protobuf:"bytes,2,rep,name=intention_list,json=intentionList,proto3" json:"intention_list,omitempty"`
	ChallengeList             []*GCGDuelChallenge      `protobuf:"bytes,1617,rep,name=challenge_list,json=challengeList,proto3" json:"challenge_list,omitempty"`
	HistoryCardList           []*GCGCard               `protobuf:"bytes,1872,rep,name=history_card_list,json=historyCardList,proto3" json:"history_card_list,omitempty"`
	Round                     uint32                   `protobuf:"varint,11,opt,name=round,proto3" json:"round,omitempty"`
	ControllerId              uint32                   `protobuf:"varint,12,opt,name=controller_id,json=controllerId,proto3" json:"controller_id,omitempty"`
	HistoryMsgPackList        []*GCGMessagePack        `protobuf:"bytes,797,rep,name=history_msg_pack_list,json=historyMsgPackList,proto3" json:"history_msg_pack_list,omitempty"`
	Unk3300_JHDDNKFPINA       uint32                   `protobuf:"varint,10,opt,name=Unk3300_JHDDNKFPINA,json=Unk3300JHDDNKFPINA,proto3" json:"Unk3300_JHDDNKFPINA,omitempty"`
	CardIdList                []uint32                 `protobuf:"varint,6,rep,packed,name=card_id_list,json=cardIdList,proto3" json:"card_id_list,omitempty"`
	Unk3300_JBBMBKGOONO       uint32                   `protobuf:"varint,15,opt,name=Unk3300_JBBMBKGOONO,json=Unk3300JBBMBKGOONO,proto3" json:"Unk3300_JBBMBKGOONO,omitempty"`
	Phase                     *GCGPhase                `protobuf:"bytes,14,opt,name=phase,proto3" json:"phase,omitempty"`
}

func (x *GCGDuel) Reset() {
	*x = GCGDuel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GCGDuel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GCGDuel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GCGDuel) ProtoMessage() {}

func (x *GCGDuel) ProtoReflect() protoreflect.Message {
	mi := &file_GCGDuel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GCGDuel.ProtoReflect.Descriptor instead.
func (*GCGDuel) Descriptor() ([]byte, []int) {
	return file_GCGDuel_proto_rawDescGZIP(), []int{0}
}

func (x *GCGDuel) GetServerSeq() uint32 {
	if x != nil {
		return x.ServerSeq
	}
	return 0
}

func (x *GCGDuel) GetShowInfoList() []*GCGControllerShowInfo {
	if x != nil {
		return x.ShowInfoList
	}
	return nil
}

func (x *GCGDuel) GetForbidFinishChallengeList() []uint32 {
	if x != nil {
		return x.ForbidFinishChallengeList
	}
	return nil
}

func (x *GCGDuel) GetCardList() []*GCGCard {
	if x != nil {
		return x.CardList
	}
	return nil
}

func (x *GCGDuel) GetUnk3300_BIANMOPDEHO() uint32 {
	if x != nil {
		return x.Unk3300_BIANMOPDEHO
	}
	return 0
}

func (x *GCGDuel) GetCostRevise() *GCGCostReviseInfo {
	if x != nil {
		return x.CostRevise
	}
	return nil
}

func (x *GCGDuel) GetGameId() uint32 {
	if x != nil {
		return x.GameId
	}
	return 0
}

func (x *GCGDuel) GetFieldList() []*GCGPlayerField {
	if x != nil {
		return x.FieldList
	}
	return nil
}

func (x *GCGDuel) GetUnk3300_CDCMBOKBLAK() []*Unk3300_ADHENCIFKNI {
	if x != nil {
		return x.Unk3300_CDCMBOKBLAK
	}
	return nil
}

func (x *GCGDuel) GetBusinessType() GCGGameBusinessType {
	if x != nil {
		return x.BusinessType
	}
	return GCGGameBusinessType_GCG_GAME_BUSINESS_TYPE_NONE
}

func (x *GCGDuel) GetIntentionList() []*GCGPVEIntention {
	if x != nil {
		return x.IntentionList
	}
	return nil
}

func (x *GCGDuel) GetChallengeList() []*GCGDuelChallenge {
	if x != nil {
		return x.ChallengeList
	}
	return nil
}

func (x *GCGDuel) GetHistoryCardList() []*GCGCard {
	if x != nil {
		return x.HistoryCardList
	}
	return nil
}

func (x *GCGDuel) GetRound() uint32 {
	if x != nil {
		return x.Round
	}
	return 0
}

func (x *GCGDuel) GetControllerId() uint32 {
	if x != nil {
		return x.ControllerId
	}
	return 0
}

func (x *GCGDuel) GetHistoryMsgPackList() []*GCGMessagePack {
	if x != nil {
		return x.HistoryMsgPackList
	}
	return nil
}

func (x *GCGDuel) GetUnk3300_JHDDNKFPINA() uint32 {
	if x != nil {
		return x.Unk3300_JHDDNKFPINA
	}
	return 0
}

func (x *GCGDuel) GetCardIdList() []uint32 {
	if x != nil {
		return x.CardIdList
	}
	return nil
}

func (x *GCGDuel) GetUnk3300_JBBMBKGOONO() uint32 {
	if x != nil {
		return x.Unk3300_JBBMBKGOONO
	}
	return 0
}

func (x *GCGDuel) GetPhase() *GCGPhase {
	if x != nil {
		return x.Phase
	}
	return nil
}

var File_GCGDuel_proto protoreflect.FileDescriptor

var file_GCGDuel_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x47, 0x43, 0x47, 0x44, 0x75, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0d, 0x47, 0x43, 0x47, 0x43, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x47, 0x43, 0x47, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x68, 0x6f,
	0x77, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x47, 0x43, 0x47,
	0x43, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x47, 0x43, 0x47, 0x44, 0x75, 0x65, 0x6c, 0x43, 0x68, 0x61,
	0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x47, 0x43,
	0x47, 0x47, 0x61, 0x6d, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x47, 0x43, 0x47, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x47,
	0x43, 0x47, 0x50, 0x56, 0x45, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x47, 0x43, 0x47, 0x50, 0x68, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x47, 0x43, 0x47, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b, 0x33,
	0x33, 0x30, 0x30, 0x5f, 0x41, 0x44, 0x48, 0x45, 0x4e, 0x43, 0x49, 0x46, 0x4b, 0x4e, 0x49, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x07, 0x0a, 0x07, 0x47, 0x43, 0x47, 0x44, 0x75, 0x65,
	0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x71, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x65, 0x71,
	0x12, 0x3c, 0x0a, 0x0e, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x47, 0x43, 0x47, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0c, 0x73, 0x68, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x40,
	0x0a, 0x1c, 0x66, 0x6f, 0x72, 0x62, 0x69, 0x64, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f,
	0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0xc0,
	0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x19, 0x66, 0x6f, 0x72, 0x62, 0x69, 0x64, 0x46, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x25, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x47, 0x43, 0x47, 0x43, 0x61, 0x72, 0x64, 0x52, 0x08, 0x63,
	0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33,
	0x30, 0x30, 0x5f, 0x42, 0x49, 0x41, 0x4e, 0x4d, 0x4f, 0x50, 0x44, 0x45, 0x48, 0x4f, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x42, 0x49, 0x41,
	0x4e, 0x4d, 0x4f, 0x50, 0x44, 0x45, 0x48, 0x4f, 0x12, 0x33, 0x0a, 0x0b, 0x63, 0x6f, 0x73, 0x74,
	0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x47, 0x43, 0x47, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x73, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0a, 0x63, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x73, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x47, 0x43, 0x47,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x09, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30,
	0x30, 0x5f, 0x43, 0x44, 0x43, 0x4d, 0x42, 0x4f, 0x4b, 0x42, 0x4c, 0x41, 0x4b, 0x18, 0xc3, 0x0f,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f, 0x41,
	0x44, 0x48, 0x45, 0x4e, 0x43, 0x49, 0x46, 0x4b, 0x4e, 0x49, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33,
	0x33, 0x30, 0x30, 0x43, 0x44, 0x43, 0x4d, 0x42, 0x4f, 0x4b, 0x42, 0x4c, 0x41, 0x4b, 0x12, 0x39,
	0x0a, 0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x47, 0x43, 0x47, 0x47, 0x61, 0x6d, 0x65, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x37, 0x0a, 0x0e, 0x69, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x47, 0x43, 0x47, 0x50, 0x56, 0x45, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x39, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0xd1, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x47, 0x43,
	0x47, 0x44, 0x75, 0x65, 0x6c, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x0d,
	0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a,
	0x11, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0xd0, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x47, 0x43, 0x47, 0x43,
	0x61, 0x72, 0x64, 0x52, 0x0f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x61, 0x72, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x43, 0x0a, 0x15, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x70,
	0x61, 0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x9d, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x47, 0x43, 0x47, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x61, 0x63, 0x6b,
	0x52, 0x12, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4d, 0x73, 0x67, 0x50, 0x61, 0x63, 0x6b,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x5f,
	0x4a, 0x48, 0x44, 0x44, 0x4e, 0x4b, 0x46, 0x50, 0x49, 0x4e, 0x41, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x4a, 0x48, 0x44, 0x44, 0x4e, 0x4b,
	0x46, 0x50, 0x49, 0x4e, 0x41, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x61, 0x72,
	0x64, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x33,
	0x30, 0x30, 0x5f, 0x4a, 0x42, 0x42, 0x4d, 0x42, 0x4b, 0x47, 0x4f, 0x4f, 0x4e, 0x4f, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x33, 0x30, 0x30, 0x4a, 0x42, 0x42,
	0x4d, 0x42, 0x4b, 0x47, 0x4f, 0x4f, 0x4e, 0x4f, 0x12, 0x1f, 0x0a, 0x05, 0x70, 0x68, 0x61, 0x73,
	0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x47, 0x43, 0x47, 0x50, 0x68, 0x61,
	0x73, 0x65, 0x52, 0x05, 0x70, 0x68, 0x61, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GCGDuel_proto_rawDescOnce sync.Once
	file_GCGDuel_proto_rawDescData = file_GCGDuel_proto_rawDesc
)

func file_GCGDuel_proto_rawDescGZIP() []byte {
	file_GCGDuel_proto_rawDescOnce.Do(func() {
		file_GCGDuel_proto_rawDescData = protoimpl.X.CompressGZIP(file_GCGDuel_proto_rawDescData)
	})
	return file_GCGDuel_proto_rawDescData
}

var file_GCGDuel_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GCGDuel_proto_goTypes = []interface{}{
	(*GCGDuel)(nil),               // 0: GCGDuel
	(*GCGControllerShowInfo)(nil), // 1: GCGControllerShowInfo
	(*GCGCard)(nil),               // 2: GCGCard
	(*GCGCostReviseInfo)(nil),     // 3: GCGCostReviseInfo
	(*GCGPlayerField)(nil),        // 4: GCGPlayerField
	(*Unk3300_ADHENCIFKNI)(nil),   // 5: Unk3300_ADHENCIFKNI
	(GCGGameBusinessType)(0),      // 6: GCGGameBusinessType
	(*GCGPVEIntention)(nil),       // 7: GCGPVEIntention
	(*GCGDuelChallenge)(nil),      // 8: GCGDuelChallenge
	(*GCGMessagePack)(nil),        // 9: GCGMessagePack
	(*GCGPhase)(nil),              // 10: GCGPhase
}
var file_GCGDuel_proto_depIdxs = []int32{
	1,  // 0: GCGDuel.show_info_list:type_name -> GCGControllerShowInfo
	2,  // 1: GCGDuel.card_list:type_name -> GCGCard
	3,  // 2: GCGDuel.cost_revise:type_name -> GCGCostReviseInfo
	4,  // 3: GCGDuel.field_list:type_name -> GCGPlayerField
	5,  // 4: GCGDuel.Unk3300_CDCMBOKBLAK:type_name -> Unk3300_ADHENCIFKNI
	6,  // 5: GCGDuel.business_type:type_name -> GCGGameBusinessType
	7,  // 6: GCGDuel.intention_list:type_name -> GCGPVEIntention
	8,  // 7: GCGDuel.challenge_list:type_name -> GCGDuelChallenge
	2,  // 8: GCGDuel.history_card_list:type_name -> GCGCard
	9,  // 9: GCGDuel.history_msg_pack_list:type_name -> GCGMessagePack
	10, // 10: GCGDuel.phase:type_name -> GCGPhase
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_GCGDuel_proto_init() }
func file_GCGDuel_proto_init() {
	if File_GCGDuel_proto != nil {
		return
	}
	file_GCGCard_proto_init()
	file_GCGControllerShowInfo_proto_init()
	file_GCGCostReviseInfo_proto_init()
	file_GCGDuelChallenge_proto_init()
	file_GCGGameBusinessType_proto_init()
	file_GCGMessagePack_proto_init()
	file_GCGPVEIntention_proto_init()
	file_GCGPhase_proto_init()
	file_GCGPlayerField_proto_init()
	file_Unk3300_ADHENCIFKNI_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GCGDuel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GCGDuel); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_GCGDuel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GCGDuel_proto_goTypes,
		DependencyIndexes: file_GCGDuel_proto_depIdxs,
		MessageInfos:      file_GCGDuel_proto_msgTypes,
	}.Build()
	File_GCGDuel_proto = out.File
	file_GCGDuel_proto_rawDesc = nil
	file_GCGDuel_proto_goTypes = nil
	file_GCGDuel_proto_depIdxs = nil
}
