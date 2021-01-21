/**
 * Copyright (c) 2021 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License-AGPL.txt in the project root for license information.
 */

import { BlobServiceClient } from "@gitpod/content-service/lib/blobs_grpc_pb";
import { BlobRef, UploadUrlRequest, UploadUrlResponse, DownloadUrlResponse, DownloadUrlRequest } from "@gitpod/content-service/lib/blobs_pb";
import { User } from "@gitpod/gitpod-protocol";
import { log } from "@gitpod/gitpod-protocol/lib/util/logging";
import * as express from 'express';
import * as grpc from "grpc";
import { injectable } from "inversify";

@injectable()
export class BlobsController {
    get apiRouter(): express.Router {
        const router = express.Router();

        router.post("/upload-url/:blobname(*)", async (req: express.Request, res: express.Response, _next: express.NextFunction) => {
            if (!req.isAuthenticated() || !User.is(req.user)) {
                res.sendStatus(500);
                return;
            }

            const blobRef = new BlobRef();
            blobRef.setName(req.params.blobname);
            blobRef.setOwnerId(req.user.id);

            const uploadUrlRequest = new UploadUrlRequest();
            uploadUrlRequest.setRef(blobRef);
            uploadUrlRequest.setMediaType(req.params.mediatype || "application/octet-stream");
            uploadUrlRequest.setTimeToLive(req.params.typetolive);

            const client = new BlobServiceClient("content-service:8080", grpc.credentials.createInsecure());
            const uploadUrlPromise = new Promise<UploadUrlResponse>((resolve, reject) => {
                client.uploadUrl(uploadUrlRequest, new grpc.Metadata(), (err: any, resp: UploadUrlResponse) => {
                    if (err) {
                        reject(err);
                    } else {
                        resolve(resp);
                    }
                });
            });
            try {
                const resp = (await uploadUrlPromise).toObject();
                res.send(`${JSON.stringify(resp)}`);
            } catch (err) {
                log.error("bloberror", err);
                res.send("error: " + err);
            } finally {
                client.close();
            }
        });

        router.get("/download-url/:blobname(*)", async (req: express.Request, res: express.Response, _next: express.NextFunction) => {
            if (!req.isAuthenticated() || !User.is(req.user)) {
                res.sendStatus(500);
                return;
            }

            log.debug(`blobs-controller: download-url '${req.params.blobname}'`)

            const blobRef = new BlobRef();
            blobRef.setName(req.params.blobname);
            blobRef.setOwnerId(req.user.id);

            const downloadUrlRequest = new DownloadUrlRequest();
            downloadUrlRequest.setRef(blobRef);

            const client = new BlobServiceClient("content-service:8080", grpc.credentials.createInsecure());
            const downloadUrlPromise = new Promise<DownloadUrlResponse>((resolve, reject) => {
                client.downloadUrl(downloadUrlRequest, new grpc.Metadata(), (err: any, resp: DownloadUrlResponse) => {
                    if (err) {
                        reject(err);
                    } else {
                        resolve(resp);
                    }
                });
            });
            try {
                const resp = (await downloadUrlPromise).toObject();
                res.send(`${JSON.stringify(resp)}`);
            } catch (err) {
                log.error("bloberror", err);
                res.send("error: " + err);
            } finally {
                client.close();
            }
        });

        return router;
    }

}