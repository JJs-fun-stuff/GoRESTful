const gulp = require("gulp");
const shell = require("gulp-shell")

gulp.task("install-binary", shell.task(['go install github.com/JJs-fun-stuff/Chap1-GettingStart/romanserver']))

gulp.task('watch', () =>{
    gulp.watch("*", ['install-binary', 'restart-supervisor'])
})

gulp.task('default'['watch'])