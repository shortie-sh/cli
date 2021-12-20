import { Arguments } from 'https://deno.land/x/yargs/deno-types.ts'


export default async function getUrl(argv: Arguments) {
    const queryStr = JSON.stringify({
        query: `query {
            getRedirect(ending: "${argv.ending}") {
              url
            }
          }`
    })
    const request = await fetch('http://localhost:3000/api/graphql', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: queryStr,
    })
    const response = await request.json()

    if(response.errors) {
        if(response.errors[0].extensions.code == "USER_INPUT_ERROR") {
            console.log(response.errors[0].message)
        } else {
            throw new Error(response.errors[0].message)
        }
    } else if(response.data.getRedirect == null) {
        console.info("No URL found")
    } else {
        console.info(response.data.getRedirect.url)
    }
}