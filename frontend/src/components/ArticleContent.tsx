import ReactMarkdown from "react-markdown";

export default function ArticleContent({ text }: { text: string }) {
  return (
    <ReactMarkdown className="p-4 overflow-y-scroll max-w-none col-span-4 prose prose-invert">
      {text}
    </ReactMarkdown>
  );
}
