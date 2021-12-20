import yargs from 'https://deno.land/x/yargs/deno.ts'
import { Arguments } from 'https://deno.land/x/yargs/deno-types.ts'
import getUrl from './getUrl.ts';
import newUrl from './newUrl.ts';

yargs(Deno.args)
.scriptName("shortie")
  .command(['$0 <url>', 'new'], 'create a new redirect', (yargs: any) => {
    yargs.positional('url', {
      describe: 'target URL to be redirected',
      type: 'string'
    });
    yargs.option('e', {
      alias: 'ending',
      describe: 'custom ending for redirect',
      type: 'string'
    })
  }, (argv: Arguments) => {
   newUrl(argv) 
  })
  .command('get <ending>', 'get existing redirect', (yargs: any) => {
      yargs.positional('ending', {
        describe: 'fetch redirect URL from ending',
        type: 'string'
      });
  }, async (argv: Arguments) => {
      await getUrl(argv)
  })
  .option("i", {
      alias: 'instance',
      describe: 'instance of shortie.sh to use',
      type: 'string',
      default: 'https://shortie.sh'
  })
  .strict()
  .demandCommand(1)
  .help()
  .parse()