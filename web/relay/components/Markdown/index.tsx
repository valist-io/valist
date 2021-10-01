import ReactMarkdown from 'react-markdown';
import remarkGfm from 'remark-gfm';

interface MarkdownProps {
  markdown: string,
}

export default function RenderMarkdown(props: MarkdownProps): JSX.Element {
  return (
    <div className="markdown">
      <ReactMarkdown skipHtml={true} remarkPlugins={[remarkGfm]}>
        {props.markdown}
      </ReactMarkdown>
    </div>
  );
}
