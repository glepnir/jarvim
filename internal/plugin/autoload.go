// Package plugin provides ...
package plugin

const AutoloadLoadEnv = `
" Load Env file and return env content
function! initself#load_env()
  let l:env_file = getenv("HOME")."/.env"
  let l:env_dict={}
  if filereadable(l:env_file)
    let l:env_content = readfile(l:env_file)
    for item in l:env_content
      let l:env_dict[split(item,"=")[0]] = split(item,"=")[1]
    endfor
    return l:env_dict
  else
    echo "env file doesn't exist"
  endif
endfunction

" Load database connection from env file
function! initself#load_db_from_env()
  let l:env = initself#load_env()
  let l:dbs={}
  for key in keys(l:env)
    if stridx(key,"DB_CONNECTION_") >= 0
      let l:dbs[split(key,"_")[2]] = l:env[key]
    endif
  endfor
  if empty(l:dbs)
    echo "Env Database config error"
  endif
  return l:dbs
endfunction
`
