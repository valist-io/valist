/* eslint-disable max-len */
import Markdown from '../Markdown';

interface RepoReadmeProps {
  repoReadme: string
}

export default function RepoReadme(props: RepoReadmeProps): JSX.Element {
  const defaultContent = `
# Readme not found
`;

  return (
    <div className="p-10">
      <Markdown markdown={props.repoReadme || defaultContent} />
    </div>
  );
}
