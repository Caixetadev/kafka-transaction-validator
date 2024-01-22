import http from "k6/http";

export const options = {
  vus: 25,
  duration: "15s",
};

export default function () {
  const payload = JSON.stringify({
    accountExternalIdDebit: "Guid",
    accountExternalIdCredit: "Guid",
    tranferTypeId: 1,
    value: Math.floor(Math.random() * 2000),
  });

  const headers = { "Content-Type": "application/json" };
  http.post("http://localhost:8080/transaction", payload, {
    headers,
  });
}
