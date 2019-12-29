filelist = getArgument();
file = split(filelist,'#');

open(file[0]);
run("Invert");
saveAs("Tiff", file[1]);
close();
run("Quit");