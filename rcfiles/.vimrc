syntax on
set number
set numberwidth=5
set nocompatible
set autoindent
set smartindent
set incsearch
set ruler
set ignorecase
set hlsearch
set vb 
set tabstop=4
set shiftwidth=4
" Press F8 to call aspell
map <F8> :w!<CR>:!aspell check %<CR>:e! %<CR>
