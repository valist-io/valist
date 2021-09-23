import ReactMarkdown from 'react-markdown';

interface MarkdownProps {
  markdown: string,
}

export default function RenderMarkdown(props: MarkdownProps): JSX.Element {
  return (
    <div className="markdown">
      <ReactMarkdown skipHtml={true}>
        {props.markdown}
      </ReactMarkdown>
    </div>
  );
}
