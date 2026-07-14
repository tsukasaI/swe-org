# /// script
# requires-python = ">=3.10"
# dependencies = [
#   "scikit-learn",
#   "numpy",
#   "pandas",
# ]
# ///

import numpy as np
import pandas as pd
from sklearn.datasets import fetch_california_housing
from sklearn.model_selection import train_test_split
from sklearn.linear_model import LinearRegression
from sklearn.metrics import mean_squared_error, mean_absolute_error, r2_score

housing = fetch_california_housing(as_frame=True)
df = housing.frame

print("=== データ概要 ===")
print(f"サンプル数: {len(df)}, 特徴量数: {len(housing.feature_names)}")
print(f"目的変数: MedHouseVal (住宅価格の中央値, 単位: $100,000)")
print()
print(housing.DESCR[:500])
print()

# ---- 特徴量の確認 ----
print("=== 特徴量一覧 ===")
print(df.describe().T.to_string())
print()

# ---- 相関係数で特徴量の関連性を確認 ----
print("=== 目的変数との相関係数 ===")
corr = df.corr()["MedHouseVal"].drop("MedHouseVal").sort_values(ascending=False)
print(corr.to_string())
print()

# ---- 特徴量選択の判断 ----
# 全特徴量を使う (ベースライン)
# California Housing は 8 特徴量しかなく、
# 多重共線性が深刻なペアも少ないため全投入が妥当
selected = housing.feature_names
print(f"=== 使用する特徴量 ({len(selected)}個, 全特徴量) ===")
for name in selected:
    print(f"  - {name}")
print()

X = df[selected]
y = df["MedHouseVal"]

X_train, X_test, y_train, y_test = train_test_split(
    X, y, test_size=0.2, random_state=42
)
print(f"訓練データ: {len(X_train)}個 / テストデータ: {len(X_test)}個")
print()

# ---- 線形回帰 ----
model = LinearRegression()
model.fit(X_train, y_train)
y_pred = model.predict(X_test)

rmse = np.sqrt(mean_squared_error(y_test, y_pred))
mae = mean_absolute_error(y_test, y_pred)
r2 = r2_score(y_test, y_pred)

print("=== 評価指標 (線形回帰) ===")
print(f"  RMSE : {rmse:.4f}  (≈ ${rmse * 100_000:,.0f})")
print(f"  MAE  : {mae:.4f}  (≈ ${mae * 100_000:,.0f})")
print(f"  R²   : {r2:.4f}")
print()

# ---- 係数の確認 ----
print("=== 回帰係数 ===")
coef_df = pd.DataFrame({
    "feature": selected,
    "coefficient": model.coef_,
}).sort_values("coefficient", key=abs, ascending=False)
print(coef_df.to_string(index=False))
print(f"\n切片 (intercept): {model.intercept_:.4f}")
print()

# ---- 考察 ----
print("=== 考察 ===")
print(f"R² = {r2:.3f} → 分散の約{r2*100:.0f}%を説明できている")
print(f"RMSE ≈ ${rmse * 100_000:,.0f} の予測誤差")
print()
print("線形回帰の限界:")
print("  - 住宅価格と特徴量の関係は非線形な部分が多い")
print("  - MedInc (所得) が支配的で、地理情報 (Lat/Long) の非線形効果を捉えられない")
print("  - 改善案: Ridge/Lasso, Random Forest, Gradient Boosting など")
