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
* PluginMount plugin mount
*/
export class PluginMount {
    /**
    * description
    */
    'description': string;
    /**
    * destination
    */
    'destination': string;
    /**
    * name
    */
    'name': string;
    /**
    * options
    */
    'options': Array<string>;
    /**
    * settable
    */
    'settable': Array<string>;
    /**
    * source
    */
    'source': string;
    /**
    * type
    */
    'type': string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "description",
            "baseName": "Description",
            "type": "string"
        },
        {
            "name": "destination",
            "baseName": "Destination",
            "type": "string"
        },
        {
            "name": "name",
            "baseName": "Name",
            "type": "string"
        },
        {
            "name": "options",
            "baseName": "Options",
            "type": "Array<string>"
        },
        {
            "name": "settable",
            "baseName": "Settable",
            "type": "Array<string>"
        },
        {
            "name": "source",
            "baseName": "Source",
            "type": "string"
        },
        {
            "name": "type",
            "baseName": "Type",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return PluginMount.attributeTypeMap;
    }
}

