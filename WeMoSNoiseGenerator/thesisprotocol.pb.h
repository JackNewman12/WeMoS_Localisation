/* Automatically generated nanopb header */
/* Generated by nanopb-0.3.5 at Fri Jul 08 16:54:31 2016. */

#ifndef PB_THESISPROTOCOL_PB_H_INCLUDED
#define PB_THESISPROTOCOL_PB_H_INCLUDED
#include "pb.h"

/* @@protoc_insertion_point(includes) */
#if PB_PROTO_HEADER_VERSION != 30
#error Regenerate this file with the current version of nanopb generator.
#endif

#ifdef __cplusplus
extern "C" {
#endif

/* Struct definitions */
typedef struct _ReceivedMessage_Device {
    pb_callback_t mac;
    pb_callback_t datapoints;
/* @@protoc_insertion_point(struct:ReceivedMessage_Device) */
} ReceivedMessage_Device;

typedef struct _ReceivedMessage {
    int32_t deviceX;
    int32_t deviceY;
    bool has_deviceVoltage;
    int32_t deviceVoltage;
    bool has_lastPost;
    bool lastPost;
    bool has_submitTime;
    int32_t submitTime;
    pb_callback_t devices;
/* @@protoc_insertion_point(struct:ReceivedMessage) */
} ReceivedMessage;

typedef struct _ReceivedMessage_DataPoint {
    int32_t rssi;
    int32_t secondsSince;
/* @@protoc_insertion_point(struct:ReceivedMessage_DataPoint) */
} ReceivedMessage_DataPoint;

/* Default values for struct fields */

/* Initializer values for message structs */
#define ReceivedMessage_init_default             {0, 0, false, 0, false, 0, false, 0, {{NULL}, NULL}}
#define ReceivedMessage_Device_init_default      {{{NULL}, NULL}, {{NULL}, NULL}}
#define ReceivedMessage_DataPoint_init_default   {0, 0}
#define ReceivedMessage_init_zero                {0, 0, false, 0, false, 0, false, 0, {{NULL}, NULL}}
#define ReceivedMessage_Device_init_zero         {{{NULL}, NULL}, {{NULL}, NULL}}
#define ReceivedMessage_DataPoint_init_zero      {0, 0}

/* Field tags (for use in manual encoding/decoding) */
#define ReceivedMessage_Device_mac_tag           1
#define ReceivedMessage_Device_datapoints_tag    2
#define ReceivedMessage_deviceX_tag              1
#define ReceivedMessage_deviceY_tag              2
#define ReceivedMessage_deviceVoltage_tag        3
#define ReceivedMessage_lastPost_tag             4
#define ReceivedMessage_submitTime_tag           5
#define ReceivedMessage_devices_tag              6
#define ReceivedMessage_DataPoint_rssi_tag       1
#define ReceivedMessage_DataPoint_secondsSince_tag 2

/* Struct field encoding specification for nanopb */
extern const pb_field_t ReceivedMessage_fields[7];
extern const pb_field_t ReceivedMessage_Device_fields[3];
extern const pb_field_t ReceivedMessage_DataPoint_fields[3];

/* Maximum encoded size of messages (where known) */
#define ReceivedMessage_DataPoint_size           17

/* Message IDs (where set with "msgid" option) */
#ifdef PB_MSGID

#define THESISPROTOCOL_MESSAGES \


#endif

#ifdef __cplusplus
} /* extern "C" */
#endif
/* @@protoc_insertion_point(eof) */

#endif
