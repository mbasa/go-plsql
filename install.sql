create or replace function get_arg_text(text,int,float) 
  returns text as :MOD, 'getArgText' 
  language c strict;
