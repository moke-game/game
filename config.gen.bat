set WORKSPACE=.\

set GEN_CLIENT=%WORKSPACE%\luban\Luban.ClientServer\Luban.ClientServer.exe
set CONF_ROOT=..\v5plan\

%GEN_CLIENT%  -j cfg --^
 -d %CONF_ROOT%\Defines\__root__.xml ^
 --input_data_dir %CONF_ROOT%\Defines ^
 --output_code_dir %WORKSPACE%/configs/code ^
 --output_data_dir %WORKSPACE%/configs/data ^
 --gen_types code_go_json,data_json ^
 --go:bright_module_name "v5game" ^
 -s server

pause