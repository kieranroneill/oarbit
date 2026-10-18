function quote(filename) {
  return `"${filename}"`;
}

export default {
  // format javascript related files
  '**/*.{cjs,js,json,mjs,ts,tsx}': (filenames) => [`prettier --write ${filenames.map(quote).join(' ')}`],
  // format go files
  '**/*.go': (filenames) => [`gofmt -w ${filenames.map(quote).join(' ')}`],
};
