/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { PlacesList } from '../models/PlacesList';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class DefaultService {

    /**
     * Get establishment information
     * @param filter Hotels, Restaurants, or Attractions
     * @param location City Name
     * @param distance
     * @returns PlacesList OK
     * @throws ApiError
     */
    public static search(
        filter: string,
        location: string,
        distance: number,
    ): CancelablePromise<PlacesList> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/search',
            query: {
                'filter': filter,
                'location': location,
                'distance': distance,
            },
            errors: {
                400: `Bad Request`,
                418: `I'm a teapot`,
                500: `Internal Server Error`,
                502: `Bad Gateway`,
            },
        });
    }

}
