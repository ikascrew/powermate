PowerMate をLinuxで使用する為のソフトです

https://github.com/awly/pmd

を元に作成しましたが、
音量を調整するような作りになっていましたので、関数を登録しておいて、
それで実行するような形式にしたものです。


udev を利用しています
上記プロジェクトと同じようにruleファイルを利用してデバイスを読み込みます

ACTION=="add", ENV{ID_USB_DRIVER}=="powermate", SYMLINK+="input/powermate", MODE="0666"


