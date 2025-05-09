const { src, dest, series } = require('gulp');
const concat = require('gulp-concat');
const terser = require('gulp-terser');
const postcss = require('gulp-postcss');
const tailwindcss = require('tailwindcss');
const autoprefixer = require('autoprefixer');
const path = require('path');

// Paths
const paths = {
  js: [
    '/app/assets/node_modules/jquery/dist/jquery.js',
    '/app/assets/node_modules/bootstrap/dist/js/bootstrap.bundle.js',
    '/app/assets/js/**/*.js'
  ],
  css: '/app/assets/css/**/*.scss',
  output: '/app/static'
};

// JavaScript Task
function jsTask() {
  return src(paths.js)
    .pipe(concat('app.js'))
    .pipe(terser())
    .pipe(dest(paths.output));
}

// CSS Task (Tailwind + Bootstrap)
function cssTask() {
  return src(paths.css)
    .pipe(postcss([
      tailwindcss('/app/assets/tailwind.config.js'),
      autoprefixer()
    ]))
    .pipe(concat('app.css'))
    .pipe(dest(paths.output));
}

exports.build = series(jsTask, cssTask);

