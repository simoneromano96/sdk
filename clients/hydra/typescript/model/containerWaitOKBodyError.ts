/**
 * ORY Hydra
 * Welcome to the ORY Hydra HTTP API documentation. You will find documentation for all HTTP APIs here.
 *
 * The version of the OpenAPI document: v1.7.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from '../api';

/**
* ContainerWaitOKBodyError container waiting error, if any
*/
export class ContainerWaitOKBodyError {
    /**
    * Details of an error
    */
    'message'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "message",
            "baseName": "Message",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return ContainerWaitOKBodyError.attributeTypeMap;
    }
}

