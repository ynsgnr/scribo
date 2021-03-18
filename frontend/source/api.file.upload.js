import * as cookie from './cookie.js'
import {IDTokenKey} from './auth.js'

var awsS3bucketName = "fileconverter";
var awsS3bucketRegion = "eu-central-1";
var awsS3IdentityPoolId = "eu-central-1:0e6ecb31-9d46-4967-b161-fb05bbee3c4f"
var awsS3IdentityPoolURL = "eu-central-1_KKbc2zHbm"

export function Upload(files){

  var file
  if (!files.length) {
    file = files
  }
  file = files[0]

  var logins = {}
  logins['cognito-idp.' + awsS3bucketRegion + '.amazonaws.com/'+ awsS3IdentityPoolURL] = cookie.getCookie(IDTokenKey)
  AWS.config.update({
    region: awsS3bucketRegion,
    credentials: new AWS.CognitoIdentityCredentials({
      IdentityPoolId: awsS3IdentityPoolId,
      Logins: logins
    })
  })

  var fileKey = AWS.config.credentials.identityId + "/" + file.name;

  // Use S3 ManagedUpload class as it supports multipart uploads
  var upload = new AWS.S3.ManagedUpload({
    params: {
      Bucket: awsS3bucketName,
      Key: fileKey,
      Body: file,
      ACL: "private"
    }
  })
  return upload.promise()
}