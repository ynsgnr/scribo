import * as cookie from './cookie.js'
import {Baseurl} from './config.js'
import * as auth from './auth.js'

const commandEndpoint = Baseurl+"/command/v1/user/:userID/command"
const getDevicesEndpoint = Baseurl+"/sync-device/v1/user/:userID/devices"

const userIDKey = ":userID"

class RateLimitException extends Error {}

export function GetDevices(){
    return get(getDevicesEndpoint)
}

export function AddDevice(deviceName, type, data){
    return sendEvent("AddDevice", JSON.stringify({
        name: deviceName,
        type: type,
        data: data
    }))
}

export function SyncRequest(deviceID, fileLocation){
    return sendEvent("Send2Device", JSON.stringify({
        deviceID: deviceID,
        fileLocation: fileLocation
    }))
}

function sendEvent(eventType,body){
    let userID = cookie.getCookie(auth.UserIDKey)
    if (userID != null){
        return fetch(commandEndpoint.replace(userIDKey,userID),
        {method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer '+ cookie.getCookie(auth.AccessTokenKey),
            'EventType': eventType,
        },
        body:body
        }).then(response=>{
            if (response.status == 429) {
                throw new RateLimitException()
            }
            if (response.status != 200){
                throw new Error()
            }
            return Promise.resolve()
        })
    }
    return Promise.resolve()
}

function get(url){
    let userID = cookie.getCookie(auth.UserIDKey)
    if (userID != null){
        return fetch(url.replace(userIDKey,userID),
        {method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer '+ cookie.getCookie(auth.AccessTokenKey),
        }}).then(response=>{
            if (response.status != 200){
                throw new Error()
            }
            return response.json()
        })
    }
    return Promise.resolve()
}