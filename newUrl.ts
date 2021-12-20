import { Arguments } from 'https://deno.land/x/yargs/deno-types.ts'
import * as validator from 'https://deno.land/x/deno_validator/mod.ts';

export default async function newUrl(argv: Arguments) {
    if(validator.isURL(argv.instance, {
        protocols: ['http','https'], 
        require_tld: false, 
        require_protocol: false, 
        require_host: true, 
        require_port: false, 
        require_valid_protocol: true, 
        allow_underscores: false, 
        host_whitelist: false, 
        host_blacklist: false, 
        allow_trailing_dot: false, 
        allow_protocol_relative_urls: false, 
        disallow_auth: false, 
        validate_length: true

    }) == false) {
        throw new Error("Invalid instance URL")
    }
    let mutationStr;
    if(argv.ending) {
        mutationStr = JSON.stringify({
            query: `mutation {
                createRedirect(url: "${argv.url}", ending: "${argv.ending}") {
                    ending
                }
            }`
        })

    } else {
        mutationStr = JSON.stringify({
            query: `mutation {
                createRedirect(url: "${argv.url}") {
                    ending
                }
            }`
        })
    } 
    const request = await fetch(`${argv.instance}/api/graphql`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: mutationStr,
    })
    const response = await request.json()
    if(response.errors) {
        if(response.errors[0].extensions.code == "USER_INPUT_ERROR") {
            console.log(response.errors[0].message)
        } else {
            throw new Error(response.errors[0].message)
        }
    } else {
        console.info(`${argv.instance}/${response.data.createRedirect.ending}`)
    }
}