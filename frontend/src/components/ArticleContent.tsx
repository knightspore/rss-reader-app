import ReactMarkdown from "react-markdown";

export default function ArticleContent({  text }: { text: string }) {

  return (
    <ReactMarkdown>{`## ${text}`}</ReactMarkdown>
  )

}
