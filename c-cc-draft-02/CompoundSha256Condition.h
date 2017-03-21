/*
 * Generated by asn1c-0.9.27 (http://lionet.info/asn1c)
 * From ASN.1 module "Crypto-Conditions"
 * 	found in "CryptoConditions.asn"
 * 	`asn1c -pdu=all -fincludes-quoted`
 */

#ifndef	_CompoundSha256Condition_H_
#define	_CompoundSha256Condition_H_


#include "asn_application.h"

/* Including external dependencies */
#include "OCTET_STRING.h"
#include "NativeInteger.h"
#include "ConditionTypes.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* CompoundSha256Condition */
typedef struct CompoundSha256Condition {
	OCTET_STRING_t	 fingerprint;
	unsigned long	 cost;
	ConditionTypes_t	 subtypes;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} CompoundSha256Condition_t;

/* Implementation */
/* extern asn_TYPE_descriptor_t asn_DEF_cost_3;	// (Use -fall-defs-global to expose) */
extern asn_TYPE_descriptor_t asn_DEF_CompoundSha256Condition;

#ifdef __cplusplus
}
#endif

#endif	/* _CompoundSha256Condition_H_ */
#include "asn_internal.h"
