# /// script
# requires-python = ">=3.10"
# dependencies = [
#   "scikit-learn",
#   "numpy",
# ]
# ///

from sklearn.datasets import load_iris
from sklearn.model_selection import train_test_split
from sklearn.linear_model import LogisticRegression
from sklearn.metrics import accuracy_score, confusion_matrix, classification_report

iris = load_iris()
X = iris.data
y = iris.target

print("特徴量：", iris.feature_names)
print("品種：", list(iris.target_names))
print("データ数：", X.shape[0], "個 / 特徴量", X.shape[1], "個")
print()

X_train, X_test, y_train, y_test = train_test_split(
        X, y, test_size=0.3, stratify=y, random_state=42
)

print(f"訓練データ： {X_train.shape[0]}個 / テストデータ: {X_test.shape[0]}個")
print()

model = LogisticRegression(max_iter=1000)
model.fit(X_train, y_train)

y_pred = model.predict(X_test)

acc = accuracy_score(y_test, y_pred)
print("=" * 30)
print(f"正解率： {acc:.3f} = {acc*100:.1f}%")
print("=" * 30)

print()

print("混合行列")
print(confusion_matrix(y_test, y_pred))
print()

print("品種ごとのレポート")
print(classification_report(y_test, y_pred, target_names=iris.target_names))
