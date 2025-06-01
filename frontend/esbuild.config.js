// frontend/esbuild.config.js
const esbuild = require('esbuild');
const path = require('path');
const fs = require('fs');

const frontendDir = __dirname;
const projectRoot = path.resolve(frontendDir, '..');
const outputDir = path.join(projectRoot, 'dist');

function cleanDirectory(dirPath) {
    if (fs.existsSync(dirPath)) {
        console.log(`Cleaning directory: ${dirPath}`);
        fs.rmSync(dirPath, { recursive: true, force: true });
    }
}

if (!fs.existsSync(outputDir)) {
    fs.mkdirSync(outputDir, { recursive: true });
}

const buildOptions = {
    entryPoints: {
        'main': path.join(frontendDir, 'src/js/main.js'),
        'template': path.join(frontendDir, 'src/ui/template.css'),
        'index': path.join(frontendDir, 'index.html'),
    },
    bundle: true,
    outdir: outputDir,
    minify: process.argv.includes('--minify'),
    sourcemap: true,

    outbase: frontendDir,

    loader: {
        '.html': 'copy',
        '.css': 'css',
        '.js': 'js',
    },
};

const isWatchMode = process.argv.includes('--watch');

cleanDirectory(outputDir);

if (isWatchMode) {
    esbuild.context(buildOptions).then(context => {
        context.watch();
        console.log('esbuild is watching for changes in ' + frontendDir + '...');
    }).catch(err => {
        console.error('esbuild watch context setup failed:', err);
        process.exit(1);
    });
} else {
    esbuild.build(buildOptions).then(result => {
        console.log('esbuild build finished successfully!');
    }).catch(err => {
        console.error('esbuild build failed:', err);
        process.exit(1);
    });
}