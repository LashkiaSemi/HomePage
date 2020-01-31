var gulp = require('gulp')
var sass = require('gulp-sass')

// 本番圧縮
gulp.task("sass", function () {
    return (
        gulp.src("src/sass/**/*.scss")
            .pipe(sass({ outputStyle: "compressed" }))
            .pipe(gulp.dest("src/css")) 
    );
});

// 作業監視用
gulp.task("sass-watch", function () {
    return gulp.watch("src/sass/**/*.scss", function () {
        return (
            gulp.src("src/sass/**/*.scss")
                .pipe(sass({ outputStyle: "expanded" }).on("error", sass.logError))
                .pipe(gulp.dest("src/css"))
        );
    });
});
