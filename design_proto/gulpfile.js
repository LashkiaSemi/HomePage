// モジュールの読み込み
var gulp = require("gulp");
var sass = require("gulp-sass");

// タスクを作成
// gulp.task("タスク名", 実行される処理)
gulp.task("sass", function () {
    return (
        gulp.src("src/sass/**/*.scss")  // 取得するファイル
            .pipe(sass({ outputStyle: "expanded" }))  // コンパイル時のオプション
            .pipe(gulp.dest("src/css"))  // 保存先
    );
});

gulp.task("sass-watch", function () {
    return gulp.watch("src/sass/**/*.scss", function () {
        return (
            gulp.src("src/sass/**/*.scss")
                .pipe(sass({ outputStyle: "expanded" }).on("error", sass.logError))
                .pipe(gulp.dest("src/css"))
        );
    });
});