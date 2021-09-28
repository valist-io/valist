export default function PublishReleaseSteps() {
  const example = `# Initialize your project
valist init
  
# Publish a release
valist publish`;
  return (
    <div className="p-8">
      <a className="float-right text-indigo-500" href="https://docs.valist.io">docs</a>
      <h2 className="text-xl">Publish a Release</h2>
      <pre className="p-4 my-4 bg-indigo-50 rounded-lg overflow-x-scroll"><code>{ example }</code></pre>
    </div>
  );
}
